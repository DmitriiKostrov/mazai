package main

import (
    "io/ioutil"
    "log"
    "os"
)

var TRIPS_FILE string = "trips.json"
//TODAY_CACHE=[]

func get_all_trips() (data string) {
    file, err := os.Open(TRIPS_FILE)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


  b, err := ioutil.ReadAll(file)
  return string(b)
}

