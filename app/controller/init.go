package controller

import (
	item_controller "flashare/app/controller/item"
	message_controller "flashare/app/controller/message"
	request_controller "flashare/app/controller/request"
	user_controller "flashare/app/controller/user"
)

type FlashareController struct {
	AuthenticationCtrl user_controller.AuthenticationController
	ItemCtrl           item_controller.ItemController
	RequestCtrl        request_controller.RequestController
	MessageCtrl        message_controller.MessageController
}

var flashareCtrl FlashareController

func GetFlashareController() FlashareController {
	return flashareCtrl
}

func InitFlashareController(authCtrl user_controller.AuthenticationController,
	itemCtrl item_controller.ItemController,
	requestCtrl request_controller.RequestController,
	messageCtrl message_controller.MessageController) {
	flashareCtrl = FlashareController{
		authCtrl,
		itemCtrl,
		requestCtrl,
		messageCtrl,
	}
}
