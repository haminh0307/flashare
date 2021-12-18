package user_repository

import (
	"flashare/entity"
)

type UserRepository interface {
	GetByEmail(email string) (entity.User, error)
	Create(entity.User) (interface{}, error)
}
