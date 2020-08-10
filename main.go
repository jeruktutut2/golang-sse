package main

import (
	"fmt"
	"log"
	"net/http"
)

var messageChan chan string

func handleSSE() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Get handshake from client...")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		messageChan = make(chan string)

		defer func() {
			close(messageChan)
			messageChan = nil
			log.Printf("client connection close")
		}()

		fluser, _ := w.(http.Flusher)

		for {

			select {
			case message := <-messageChan:
				fmt.Fprintf(w, "data: %v\n\n", message)
				fluser.Flush()
			case <-r.Context().Done():
				return
			}
		}
	}
}

func sendMessage(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if messageChan != nil {
			log.Printf("pring message to client...")
			messageChan <- message
		}
	}
}

func main() {
	http.HandleFunc("/handshake", handleSSE())
	http.HandleFunc("/sendmessage", sendMessage("ini message"))
	log.Printf("starting HTTP server...")
	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:8181", nil))
}
