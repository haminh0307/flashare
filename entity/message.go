package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender   string             `json:"sender" binding:"required"`
	Receiver string             `json:"receiver" binding:"required"`
	Content  string             `json:"content" binding:"required"`
	Time     time.Time          `bson:"time" json:"time" binding:"required"`
}
