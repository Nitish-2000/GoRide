package main

import (
	"fmt"
	"log"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"net/http"
)

var httpAddr = ":8083"

func main() {

	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)

	//Setting up HTTP Server to listen to incoming Request

	mux := http.NewServeMux()
	httpHandler := h.HTTPHandler{
		Service: svc,
	}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := http.Server{
		Addr:    fmt.Sprintf(httpAddr),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error while starting the server %v", err)
	}

}
