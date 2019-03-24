package main

import "fmt"
import "net/http"
import "strings"
import "log"

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //анализ аргументов,
    fmt.Println("form:", r.Form)  // ввод информации о форме на стороне сервера
    fmt.Println("path:", r.URL.Path)
    fmt.Println("scheme:", r.URL.Scheme)
    fmt.Println("url:", r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello World!") // отправляем данные на клиентскую сторону
}

func handleRequests() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "templates/index.html")
	}); 
	http.HandleFunc("/trips", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, get_all_trips()) 
	});
	http.HandleFunc("/trip/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "templates/new.html")
	}); 
	
    err := http.ListenAndServe(":9000", nil) // задаем слушать порт
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func main() {
	handleRequests()
}