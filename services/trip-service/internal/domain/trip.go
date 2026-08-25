package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct{
	Id primitive.ObjectID
	UserID string
	PackageSlug string //eg: Van, luxuru, Sedan
	TotalPriceInCents float64


}


type TripModel struct{
	ID   primitive.ObjectID
	UserID string
	Status string
	RideFare RideFareModel
}

type Triprepository interface{
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface{
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}