package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseAccount struct {
	ID      primitive.ObjectID
	Name    string
	Balance int
}
