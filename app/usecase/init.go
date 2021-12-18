package usecase

import (
	"flashare/app/usecase/item"
	"flashare/app/usecase/user"
)

type FlashareUsecase struct {
	AuthenticationUC user_usecase.AuthenticationUsecase
	ItemUC           item_usecase.ItemUsecase
}

var flashareUC FlashareUsecase

func GetFlashareUsecase() FlashareUsecase {
	return flashareUC
}

func InitFlashareUsecase(authUC user_usecase.AuthenticationUsecase, itemUC item_usecase.ItemUsecase) {
	flashareUC = FlashareUsecase{
		authUC,
		itemUC,
	}
}
