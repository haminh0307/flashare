package controller

import (
	item_controller "flashare/app/controller/item"
	request_controller "flashare/app/controller/request"
	user_controller "flashare/app/controller/user"
)

type FlashareController struct {
	AuthenticationCtrl user_controller.AuthenticationController
	ItemCtrl           item_controller.ItemController
	RequestCtrl        request_controller.RequestController
}

var flashareCtrl FlashareController

func GetFlashareController() FlashareController {
	return flashareCtrl
}

func InitFlashareController(authCtrl user_controller.AuthenticationController, itemCtrl item_controller.ItemController, requestCtrl request_controller.RequestController) {
	flashareCtrl = FlashareController{
		authCtrl,
		itemCtrl,
		requestCtrl,
	}
}
