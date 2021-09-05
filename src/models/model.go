package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn      string             `json:"isbn-10,omitempty" bson:"isbn,omitempty"`
	Title     string             `json:"title" bson:"title,omitempty"`
	Author    string             `json:"author" bson:"author,omitempty"`
	Create_at time.Time          `json:"create_at,omitempty"`
	Update_at time.Time          `json:"update_at,omitempty"`
}
