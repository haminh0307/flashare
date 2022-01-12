package controller

import (
	item_controller "flashare/app/controller/item"
	message_controller "flashare/app/controller/message"
	request_controller "flashare/app/controller/request"
	review_controller "flashare/app/controller/review"
	user_controller "flashare/app/controller/user"
)

type FlashareController struct {
	AuthenticationCtrl user_controller.AuthenticationController
	ProfileCtrl        user_controller.ProfileController
	ItemCtrl           item_controller.ItemController
	RequestCtrl        request_controller.RequestController
	MessageCtrl        message_controller.MessageController
	ReviewCtrl         review_controller.ReviewController
}

var flashareCtrl FlashareController

func GetFlashareController() FlashareController {
	return flashareCtrl
}

func InitFlashareController(authCtrl user_controller.AuthenticationController,
	profileCtrl user_controller.ProfileController,
	itemCtrl item_controller.ItemController,
	requestCtrl request_controller.RequestController,
	messageCtrl message_controller.MessageController,
	reviewCtrl review_controller.ReviewController) {
	flashareCtrl = FlashareController{
		authCtrl,
		profileCtrl,
		itemCtrl,
		requestCtrl,
		messageCtrl,
		reviewCtrl,
	}
}
