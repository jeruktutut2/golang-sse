package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int

func main1() {
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/sse/dashboard", dashboardHandler)
	http.HandleFunc("/sse/hit1", hit1)
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	counter++
	fmt.Fprintf(w, "data: %v\n\n", counter)
	fmt.Printf("data: %v\n", counter)
}

func hit1(w http.ResponseWriter, r *http.Request) {
	dashboardHandler(w, r)
}
