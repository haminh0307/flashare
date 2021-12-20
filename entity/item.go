package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Category    string             `bson:"category" json:"category"`
	PhotosLink  []string           `bson:"photos_link" json:"photos_link"`
	Description string             `bson:"description" json:"description"`
	DueDate     *time.Time         `bson:"due_date,omitempty" json:"due_date,omitempty"`
	UploadedBy  string             `bson:"uploaded_by" json:"uploaded_by"`
	Status      string             `bson:"status" json:"status"`
}
