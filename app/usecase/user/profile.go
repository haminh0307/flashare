package user_usecase

import (
	"flashare/entity"
)

type ProfileUsecase interface {
	Get(userId string) (entity.User, error)
	UpdateInfo(userId, fullName, phoneNumber, address string) error
	ChangePassword(userId, oldPwd, newPwd string) error
	UpdateAvatar(userId, avtLink string) error
}
