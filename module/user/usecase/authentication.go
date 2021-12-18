package user_usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"flashare/app/repository/user"
	"flashare/app/usecase/user"
	"flashare/entity"
	"flashare/errors"
)

type authenticationUsecaseImpl struct {
	repo user_repository.UserRepository
}

func NewAuthenticationUsecase(uRepo user_repository.UserRepository) user_usecase.AuthenticationUsecase {
	return &authenticationUsecaseImpl{
		uRepo,
	}
}

func (authUC *authenticationUsecaseImpl) SignIn(email, pwd string) (entity.User, error) {
	user, err := authUC.repo.GetByEmail(email)

	// email not exist
	if err == mongo.ErrNoDocuments {
		return entity.User{}, flashare_errors.ErrorInvalidCredentials
	}
	// internal server error
	if err != nil {
		return entity.User{}, flashare_errors.ErrorInternalServerError
	}
	// if password not matched
	if err := bcrypt.CompareHashAndPassword(user.PasswordHashCode, []byte(pwd)); err != nil {
		return entity.User{}, flashare_errors.ErrorInvalidCredentials
	}
	// matched
	return user, err
}

func (authUC *authenticationUsecaseImpl) SignUp(user entity.User) (primitive.ObjectID, error) {
	_, err := authUC.repo.GetByEmail(user.Email)

	// email already exists
	if err == nil {
		return primitive.ObjectID{}, flashare_errors.ErrorEmailAlreadyExists
	}
	// internal server error
	if err != mongo.ErrNoDocuments {
		return primitive.ObjectID{}, flashare_errors.ErrorInternalServerError
	}

	// store in db
	user_id, err := authUC.repo.Create(user)
	if err != nil {
		return primitive.ObjectID{}, flashare_errors.ErrorInternalServerError
	}
	return user_id.(primitive.ObjectID), err
}
