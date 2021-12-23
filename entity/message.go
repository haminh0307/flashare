package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender   string             `json:"sender"`
	Receiver string             `json:"receiver"`
	Content  string             `json:"content"`
	Time     time.Time          `bson:"time" json:"time"`
}
