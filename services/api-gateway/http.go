package main

import (
	"bytes"
	"encoding/json"

	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {

	var reqBody PreviewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Error Parsing the Response", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	//Checkingi if the userID is empty or not. If it is empty then we will return the error message to the user.
	if reqBody.UserID == "" {
		http.Error(w, "UserId is Required", http.StatusBadRequest)
		return
	}

	byte, _ := json.Marshal(r.Body)

	reader := bytes.NewReader(byte)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)

	if err != nil {
		http.Error(w, "Error while calling the trip service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var respBody any
	
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, "Error Parsing the Response", http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, respBody)

}
