package server

import (
	"net/http"
	"os"
	"time"
)

func Get() *http.Server {
	port := os.Getenv("port")
	mux := http.NewServeMux()
	SetHandlers(mux)

	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return server
}
