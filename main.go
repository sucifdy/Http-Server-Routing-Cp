package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler untuk menampilkan waktu saat ini dalam format "Weekday, Day Month Year"
func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		formattedTime := currentTime.Format("Monday, 2 January 2006")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(formattedTime))
	}
}

// Handler untuk menyapa pengguna dengan atau tanpa parameter `name`
func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		message := "Hello there"

		if name != "" {
			message = fmt.Sprintf("Hello, %s!", name)
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	}
}

func main() {
	// Routing dengan http.HandleFunc untuk masing-masing handler
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())

	// Memulai server pada localhost dengan port 8080
	http.ListenAndServe("localhost:8080", nil)
}
