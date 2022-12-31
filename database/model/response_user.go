package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseUser struct {
	ID        primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
	Accounts  []ResponseAccount
}
