package user_usecase

import (
	"flashare/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthenticationUsecase interface {
	SignIn(email, password string) (entity.User, error)
	SignUp(entity.User) (primitive.ObjectID, error)
}
