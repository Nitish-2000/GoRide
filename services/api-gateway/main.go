package main

import (
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway by Nitish_MNT")

	mux := http.NewServeMux()

	//In Older vserion of Go we used to use gorilla/mux for routing 
	// but now we can use the default http package to handle routing
	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	server := &http.Server{
		Addr:httpAddr,
		Handler: mux,
	}
	if err := server.ListenAndServe();err!=nil{
		log.Printf("Http Server Error %v ", err)
	}
}
