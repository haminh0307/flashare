package user_repository

import (
	"flashare/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Get(id primitive.ObjectID) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	Create(entity.User) (interface{}, error)
	Update(entity.User) (bool, error)
}
