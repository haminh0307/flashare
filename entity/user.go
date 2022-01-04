package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email            string             `bson:"email,omitempty" json:"email"`
	PasswordHashCode []byte             `bson:"password_hash_code,omitempty" json:"password_hash_code"`
	FullName         string             `bson:"full_name,omitempty" json:"full_name"`
	AvatarLink       string             `bson:"avatar_link,omitempty" json:"avatar_link"`
	PhoneNumber      string             `bson:"phone_number,omitempty" json:"phone_number"`
	Address          string             `bson:"address,omitempty" json:"address"`
	Rate             float64            `bson:"rate,omitempty" json:"rate"`
}
