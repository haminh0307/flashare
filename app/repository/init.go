package repository

import (
	"flashare/app/repository/item"
)

type FlashareRepo struct {
	ItemRepo item_repository.ItemRepository
}

var flashareRepo FlashareRepo

func GetFlashareRepo() FlashareRepo {
	return flashareRepo
}

func InitFlashareRepo(itemRepo item_repository.ItemRepository) {
	flashareRepo = FlashareRepo{
		itemRepo,
	}
}