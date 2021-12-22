package repository

import (
	item_repository "flashare/app/repository/item"
	request_repository "flashare/app/repository/request"
	user_repository "flashare/app/repository/user"
)

type FlashareRepo struct {
	UserRepo    user_repository.UserRepository
	ItemRepo    item_repository.ItemRepository
	RequestRepo request_repository.RequestRepository
	// UserRepo
}

var flashareRepo FlashareRepo

func GetFlashareRepo() FlashareRepo {
	return flashareRepo
}

func InitFlashareRepo(userRepo user_repository.UserRepository, itemRepo item_repository.ItemRepository, requestRepo request_repository.RequestRepository) {
	flashareRepo = FlashareRepo{
		userRepo,
		itemRepo,
		requestRepo,
	}
}
