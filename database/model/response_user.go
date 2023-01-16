package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseUser struct {
	ID       primitive.ObjectID
	Name     string
	Nick     string
	Email    string
	Accounts []ResponseAccount
}
