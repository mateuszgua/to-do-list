package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Type      string
	FirstName string
	Balance   uint
	UserId    primitive.ObjectID
}
