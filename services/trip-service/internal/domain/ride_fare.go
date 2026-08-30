package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct{
	Id primitive.ObjectID
	UserID string
	PackageSlug string //eg: Van, luxuru, Sedan
	TotalPriceInCents float64


} 