package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	STATIC_DIR = "/static/"
	PORT       = "8080"
)

var (
	tripTemplate = template.Must(template.ParseFiles("templates/trip.html"))
)

type TripPageData struct {
	ID     string
	METHOD string
}

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func GetTripPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := TripPageData{ID: vars["id"], METHOD: "POST"}
	tripTemplate.Execute(w, data)
}

func GetAllTrips(w http.ResponseWriter, r *http.Request) {
	data, err := get_all_trips()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		fmt.Fprintf(w, "%s", data)
	}
}

func EditTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("id: %s\n", vars["id"])
	//GetHomePage(w, r)
	var trip = make(map[string]interface{})

	trip["time"] = r.FormValue("time")
	trip["delay"] = r.FormValue("delay")
	trip["driver"] = r.FormValue("driver")
	trip["desc"] = r.FormValue("desc")
	trip["to"] = r.FormValue("to")
	trip["from"] = r.FormValue("from")

	var days = make([]map[string]interface{}, 7)
	for i := 0; i < 7; i++ {
		var day = make(map[string]interface{})
		day["seats"], _ = strconv.Atoi(r.FormValue(fmt.Sprintf("seats%d", i)))
		day["guests"] = 0
		days[i] = day
	}
	trip["days"] = days

	var id = r.FormValue("id")
	if id == "-1" {
		add_trip(trip)
	} else {
		var intID, _ = strconv.Atoi(id)
		edit_trip(trip, intID)
	}
}

func GetTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if id == -1 {
		fmt.Fprintf(w, get_new_trip())
	} else {
		var trip, _ = get_trip(id)
		fmt.Fprintf(w, trip)
	}
}

func DeleteTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	delete_trip(id)
}

func DeleteGuest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	day, _ := strconv.Atoi(vars["day"])
	update_guest(id, day, -1)
}

func AddGuest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	day, _ := strconv.Atoi(vars["day"])
	update_guest(id, day, 1)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Server CSS, JS & Images Statically.
	router.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	return router
}

func handleRequests() {
	var r = NewRouter()

	r.HandleFunc("/", GetHomePage).Methods("GET")
	r.HandleFunc("/trip/{id}/edit", GetTripPage).Methods("GET")

	r.HandleFunc("/trips", GetAllTrips).Methods("GET")
	r.HandleFunc("/trip/{id}", GetTrip).Methods("GET")
	r.HandleFunc("/trip/{id}", EditTrip).Methods("POST")
	r.HandleFunc("/trip/{id}", DeleteTrip).Methods("DELETE")

	r.HandleFunc("/trip/{id}/{day}/guest", AddGuest).Methods("POST")
	r.HandleFunc("/trip/{id}/{day}/guest", DeleteGuest).Methods("DELETE")

	err := http.ListenAndServe(":"+PORT, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	handleRequests()
}
