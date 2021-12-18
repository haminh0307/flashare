package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email            string             `bson:"email" json:"email"`
	PasswordHashCode []byte             `bson:"password_hash_code" json:"password_hash_code"`
	FullName         string             `bson:"full_name" json:"full_name"`
	AvatarLink       string             `bson:"avatar_link" json:"avatar_link"`
	PhoneNumber      string             `bson:"phone_number" json:"phone_number"`
	Address          string             `bson:"address" json:"address"`
	Rate             float64            `bson:"rate" json:"rate"`
}
