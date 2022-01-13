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

type profileUsecaseImpl struct {
	repo user_repository.UserRepository
}

func NewProfileUsecase(uRepo user_repository.UserRepository) user_usecase.ProfileUsecase {
	return &profileUsecaseImpl{
		uRepo,
	}
}

func (pUC *profileUsecaseImpl) Get(userId string) (entity.User, error) {
	id, err := primitive.ObjectIDFromHex(userId)

	// invalid id
	if err != nil {
		return entity.User{}, flashare_errors.ErrorInvalidObjectID
	}

	user, err := pUC.repo.Get(id)
	// user not exists
	if err == mongo.ErrNoDocuments {
		return entity.User{}, flashare_errors.ErrorUserNotExists
	}
	// internal server error
	if err != nil {
		return entity.User{}, flashare_errors.ErrorInternalServerError
	}
	// found
	return user, err
}

func (pUC *profileUsecaseImpl) UpdateInfo(userId, fullName, phoneNumber, address string) error {
	id, err := primitive.ObjectIDFromHex(userId)
	// invalid id
	if err != nil {
		return flashare_errors.ErrorInvalidObjectID
	}

	// update info
	updateUser := entity.User{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Address:     address,
	}
	updated, err := pUC.repo.Update(updateUser)

	// internal server error
	if err != nil {
		return flashare_errors.ErrorInternalServerError
	}
	// user not exists
	if !updated {
		return flashare_errors.ErrorUserNotExists
	}

	// updated successfully
	return err
}

func (pUC *profileUsecaseImpl) ChangePassword(userId, oldPwd, newPwd string) error {
	id, err := primitive.ObjectIDFromHex(userId)
	// invalid id
	if err != nil {
		return flashare_errors.ErrorInvalidObjectID
	}

	user, err := pUC.repo.Get(id)

	// user not exist
	if err == mongo.ErrNoDocuments {
		return flashare_errors.ErrorUserNotExists
	}
	// internal server error
	if err != nil {
		return flashare_errors.ErrorInternalServerError
	}
	// if password not matched
	if err := bcrypt.CompareHashAndPassword(user.PasswordHashCode, []byte(oldPwd)); err != nil {
		return flashare_errors.ErrorInvalidCredentials
	}

	// generate password hash
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.MinCost)
	if err != nil {
		return flashare_errors.ErrorInternalServerError
	}

	// update info
	updateUser := entity.User{
		ID:               id,
		PasswordHashCode: pwdHash,
	}
	updated, err := pUC.repo.Update(updateUser)

	// internal server error
	if err != nil || !updated {
		return flashare_errors.ErrorInternalServerError
	}

	// updated successfully
	return err
}

func (pUC *profileUsecaseImpl) UpdateAvatar(userId, avtLink string) error {
	id, err := primitive.ObjectIDFromHex(userId)
	// invalid id
	if err != nil {
		return flashare_errors.ErrorInvalidObjectID
	}

	// update info
	updateUser := entity.User{
		ID:         id,
		AvatarLink: avtLink,
	}
	updated, err := pUC.repo.Update(updateUser)

	// internal server error
	if err != nil {
		return flashare_errors.ErrorInternalServerError
	}
	// user not exists
	if !updated {
		return flashare_errors.ErrorUserNotExists
	}

	// updated successfully
	return err
}
