package main

import (
	"log"
	"net/http"

	BasicWebServer "github.com/NursultanNurgaliyev/BasicWebServer"
)

func main() {
	server := BasicWebServer.NewServer()

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.GetDataHandler(w, r)
		case "POST":
			server.PostDataHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/data/", server.DeleteDataHandler)
	http.HandleFunc("/stats", server.StatsHandler)

	go server.StartBackgroundWorker()

	go func() {
		log.Println("Starting server on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s", err)
		}
	}()

	server.Shutdown()
}
