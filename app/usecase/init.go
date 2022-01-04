package usecase

import (
	item_usecase "flashare/app/usecase/item"
	message_usecase "flashare/app/usecase/message"
	request_usecase "flashare/app/usecase/request"
	user_usecase "flashare/app/usecase/user"
)

type FlashareUsecase struct {
	AuthenticationUC user_usecase.AuthenticationUsecase
	ProfileUC        user_usecase.ProfileUsecase
	ItemUC           item_usecase.ItemUsecase
	RequestUC        request_usecase.RequestUsecase
	MessageUC        message_usecase.MessageUsecase
}

var flashareUC FlashareUsecase

func GetFlashareUsecase() FlashareUsecase {
	return flashareUC
}

func InitFlashareUsecase(authUC user_usecase.AuthenticationUsecase,
	profileUC user_usecase.ProfileUsecase,
	itemUC item_usecase.ItemUsecase,
	requestUC request_usecase.RequestUsecase,
	messageUC message_usecase.MessageUsecase) {
	flashareUC = FlashareUsecase{
		authUC,
		profileUC,
		itemUC,
		requestUC,
		messageUC,
	}
}
