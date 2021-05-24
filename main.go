package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("starting scraper")
	http.HandleFunc("/v0/PageDetails", pageDetails)
	log.Println("Starting API service on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("[API STARTUP ERROR] " + err.Error())
	}
}

func pageDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching page details")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	defer r.Body.Close()
	JSONVal, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		fmt.Fprint(w, `{"message": "`+err.Error()+`"}`)
	}
	response := pageDetailsHandler(string(JSONVal))
	fmt.Fprint(w, response)
}
