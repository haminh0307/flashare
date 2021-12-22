package usecase

import (
	item_usecase "flashare/app/usecase/item"
	request_usecase "flashare/app/usecase/request"
	user_usecase "flashare/app/usecase/user"
)

type FlashareUsecase struct {
	AuthenticationUC user_usecase.AuthenticationUsecase
	ItemUC           item_usecase.ItemUsecase
	RequestUC        request_usecase.RequestUsecase
}

var flashareUC FlashareUsecase

func GetFlashareUsecase() FlashareUsecase {
	return flashareUC
}

func InitFlashareUsecase(authUC user_usecase.AuthenticationUsecase, itemUC item_usecase.ItemUsecase, requestUC request_usecase.RequestUsecase) {
	flashareUC = FlashareUsecase{
		authUC,
		itemUC,
		requestUC,
	}
}
