package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMetaData struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Nick           string             `json:"user_nick,omitempty" bson:"user_nick,omitempty"`
	Password       string             `json:"user_password,omitempty" bson:"user_password,omitempty"`
	Email          string             `json:"user_email,omitempty" bson:"user_email,omitempty"`
	IndexationDate time.Time          `json:"indexation_date,omitempty" bson:"indexation_date,omitempty"`
}
