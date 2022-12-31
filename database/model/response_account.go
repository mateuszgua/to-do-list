package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseAccount struct {
	ID        primitive.ObjectID
	FirstName string
	Balance   int
}
