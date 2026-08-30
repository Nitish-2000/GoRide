package main

import "ride-sharing/shared/types"

//This file holds the data structure from the request;

type PreviewTripRequest struct {
	UserID      string `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

//`json:"userID"` This is called marshelling the data.
//  When the data received will be in the name of "userID". 
// But inside the God application we will use as "UserId"
