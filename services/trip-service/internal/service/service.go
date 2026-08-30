package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct{
    repo domain.Triprepository
}

func NewService(repo domain.Triprepository)*service{
    return &service{
        repo:repo,
    }
}

//Making this as an methode for the "Service"
func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error){
    
    t:= &domain.TripModel{
        ID:primitive.NewObjectID(),
        UserID: fare.UserID,
        Status: "Pending",
        RideFare: fare,

    }
    return s.repo.CreateTrip(ctx,t)
}
