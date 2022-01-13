package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender   string             `bson:"sender" json:"sender"`
	Receiver string             `bson:"receiver" json:"receiver"`
	Rate     int                `bson:"rate" json:"rate"`
	Review   string             `bson:"review" json:"review"`
}
