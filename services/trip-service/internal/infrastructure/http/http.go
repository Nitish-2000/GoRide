package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HTTPHandler struct {
	Service domain.TripService
}

func (s *HTTPHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {

	var reqBody PreviewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Error Parsing the Response", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	fare := &domain.RideFareModel{
		UserID: "24",
	}


	ctx := r.Context()

	t, err := s.Service.CreateTrip(ctx, fare);
	if err!=nil{
		log.Panicf("Error Creating Trip Preview %v", err)
	}

	writeJSON(w, http.StatusOK,t)


}

func writeJSON(w http.ResponseWriter, status int, data any) error{
	w.Header().Set("Content-Type", "pplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}



//This file holds the data structure from the request;

type PreviewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
