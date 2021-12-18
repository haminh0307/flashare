package repository

import (
	"flashare/app/repository/item"
	"flashare/app/repository/user"
)

type FlashareRepo struct {
	UserRepo user_repository.UserRepository
	ItemRepo item_repository.ItemRepository
	// UserRepo
}

var flashareRepo FlashareRepo

func GetFlashareRepo() FlashareRepo {
	return flashareRepo
}

func InitFlashareRepo(userRepo user_repository.UserRepository, itemRepo item_repository.ItemRepository) {
	flashareRepo = FlashareRepo{
		userRepo,
		itemRepo,
	}
}
