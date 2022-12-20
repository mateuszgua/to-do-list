package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskMetaData struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"task_name,omitempty" bson:"task_name,omitempty"`
	IndexationDate time.Time          `json:"task_date_added,omitempty" bson:"task_date_added,omitempty"`
	Completed      int                `json:"task_completed,omitempty" bson:"task_completed,omitempty"`
	CompletedDate  time.Time          `json:"task_date_completed,omitempty" bson:"task_date_completed,omitempty"`
}
