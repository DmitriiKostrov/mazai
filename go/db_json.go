package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Day struct {
	seats  int `json:"seats"`
	guests int `json:"guests"`
}

type Trip struct {
	to     string `json:"to"`
	delay  string `json:"delay"`
	desc   string `json:"desc"`
	driver string `json:"driver"`
	time   string `json:"time"`
	from   string `json:"from"`
	days   []Day  `json:"days"`
}

var TRIPS_FILE string = "trips.json"
var TRIPS_CACHE_INITED = false
var TRIPS_CACHE []map[string]interface{}

func get_all_trips() (string, error) {
	if TRIPS_CACHE_INITED {
		var bolb, err = json.Marshal(TRIPS_CACHE)
		return string(bolb), err
	}

	file, err := os.Open(TRIPS_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, &TRIPS_CACHE)
	//json.Unmarshal(b, &dat)
	//fmt.Println(dat)

	TRIPS_CACHE_INITED = true
	return string(b), err
}

func get_trip(id int) (string, error) {
	get_all_trips()
	if len(TRIPS_CACHE) <= id {
		return "", errors.New("Wrong ID")
	}
	var bolb, err = json.Marshal(TRIPS_CACHE[id])
	return string(bolb), err
}

func save_trips() {
	var bolb, _ = json.Marshal(TRIPS_CACHE)
	var err = ioutil.WriteFile(TRIPS_FILE, bolb, 0644)
	if err != nil {
		panic(err)
	}
}

func delete_trip(id int) {
	get_all_trips()
	if len(TRIPS_CACHE) <= id {
		return
	}

	TRIPS_CACHE = append(TRIPS_CACHE[:id], TRIPS_CACHE[id+1:]...)
	save_trips()
}

func update_guest(id int, day int, upd int) {
	get_all_trips()
	if len(TRIPS_CACHE) <= id {
		return
	}

	md, err := TRIPS_CACHE[id]["days"].([]interface{})
	fmt.Println(err, md, md[day])
	mg, err := md[day].(map[string]interface{})

	g := mg["guests"].(float64)
	mg["guests"] = g + (float64)(upd)
	save_trips()
}

func add_trip(trip map[string]interface{}) {
	get_all_trips()
	TRIPS_CACHE = append(TRIPS_CACHE, trip)
	save_trips()
}

func edit_trip(trip map[string]interface{}, id int) {
	get_all_trips()
	if len(TRIPS_CACHE) <= id {
		return
	}
	TRIPS_CACHE[id] = trip
	save_trips()
}

func get_new_trip() string {
	return `{
        "to": "To",
        "delay": "10",
        "desc": "Desc",
        "driver": "Driver",
        "time": "18:10",
        "from": "From",
        "days": [
            {
                "seats": 4,
                "guests": 0
            },
            {
                "seats": 4,
                "guests": 0
            },
            {
                "seats": 4,
                "guests": 0
            },
            {
                "seats": 4,
                "guests": 0
            },
            {
                "seats": 4,
                "guests": 0
            },
            {
                "seats": 0,
                "guests": 0
            },
            {
                "seats": 0,
                "guests": 0
            }
        ]
    }`
}
