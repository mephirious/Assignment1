package BasicWebServer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	mu         sync.Mutex
	data       map[string]string
	requests   int
	shutdownCh chan struct{}
}

func NewServer() *Server {
	return &Server{
		data:       make(map[string]string),
		shutdownCh: make(chan struct{}),
	}
}

func (s *Server) PostDataHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch r.Method {
	case "POST":
		var d map[string]string
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, "Unexpected error occurred while decoding JSON request", http.StatusBadRequest)
			slog.Error("Error: Failed to decode request")
			return
		}
		for key, value := range d {
			s.data[key] = value
			slog.Info(fmt.Sprintf("Data with key %s and value %s added successfully", key, value))
		}
		s.requests++
		slog.Info("POST request for data finished successfully")
	default:
		http.Error(w, "Invalid method for this endpoint", http.StatusMethodNotAllowed)
		slog.Error("Error: Invalid method")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetDataHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch r.Method {
	case "GET":
		var message []Data
		for key, value := range s.data {
			message = append(message, Data{Key: key, Value: value})
		}
		err := json.NewEncoder(w).Encode(message)
		s.requests++
		slog.Info("GET request for data finished successfully")
		if err != nil {
			http.Error(w, "Failed to encode data", http.StatusInternalServerError)
			slog.Error("Error: Failed to encode data")
			return
		}
	default:
		http.Error(w, "Invalid method for this endpoint", http.StatusMethodNotAllowed)
		slog.Error("Error: Invalid method")
		return
	}
}

func (s *Server) StatsHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch r.Method {
	case "GET":
		s.requests++
		err := json.NewEncoder(w).Encode(Stats{Requests: s.requests})
		if err != nil {
			http.Error(w, "Failed to encode stats", http.StatusInternalServerError)
			slog.Error("Error: Failed to encode stats")
			return
		}
		slog.Info("GET request for stats finished successfully")
	default:
		http.Error(w, "Invalid method for this endpoint", http.StatusMethodNotAllowed)
		slog.Error("Error: Invalid method")
		return
	}
}

func (s *Server) DeleteDataHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch r.Method {
	case "DELETE":
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			slog.Error("Error: Invalid URL format")
			return
		}

		key := parts[2]
		_, isExist := s.data[key]
		if key == "" || !isExist {
			http.Error(w, "Key not found", http.StatusNotFound)
			slog.Error("Error: Key not found")
			return
		}
		delete(s.data, key)
		slog.Info("DELETE request for data finished successfully")
	default:
		http.Error(w, "Invalid method for this endpoint", http.StatusMethodNotAllowed)
		slog.Error("Error: Invalid method")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) StartBackgroundWorker() {
	for {
		select {
		case <-s.shutdownCh:
			return
		default:
			time.Sleep(5 * time.Second)
			slog.Info(fmt.Sprintf("Number of requests: %d. Size of the database: %d", s.requests, len(s.data)))
		}
	}
}

func (s *Server) Shutdown() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-signalChannel:
		log.Printf("Received signal: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server := &http.Server{Addr: ":8080"}
	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("Error shutting down server: %v", err)
	} else {
		log.Println("Server gracefully stopped.")
	}

	close(s.shutdownCh)
}
