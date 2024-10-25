package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	for {
		number := rand.Intn(100)
		fmt.Fprintf(w, "data: %d\n\n", number)

		flusher.Flush()

		time.Sleep(3 * time.Second)
	}
}

func main() {
	http.HandleFunc("/events-streaming", SSEHandler)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
