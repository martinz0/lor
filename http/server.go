package http

import (
	"context"
	"log"
	"net/http"

	"github.com/martinz0/lor/signal"
)

func ListenAndServe(addr string, handler http.Handler) {
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go listenAndServe(server)
	signal.Serve()
	shutdown(server)
}

func listenAndServe(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func shutdown(server *http.Server) {
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
}
