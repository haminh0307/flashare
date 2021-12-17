package usecase

import (
	item_usecase "flashare/app/usecase/item"
)

type FlashareUsecase struct {
	ItemUC item_usecase.ItemUsecase
}

var flashareUC FlashareUsecase

func GetFlashareUsecase() FlashareUsecase {
	return flashareUC
}

func InitFlashareUsecase(itemUC item_usecase.ItemUsecase) {
	flashareUC = FlashareUsecase{
		itemUC,
	}
}
