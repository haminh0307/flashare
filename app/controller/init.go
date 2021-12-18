package controller

import (
	"flashare/app/controller/item"
	"flashare/app/controller/user"
)

type FlashareController struct {
	AuthenticationCtrl user_controller.AuthenticationController
	ItemCtrl           item_controller.ItemController
}

var flashareCtrl FlashareController

func GetFlashareController() FlashareController {
	return flashareCtrl
}

func InitFlashareController(authCtrl user_controller.AuthenticationController, itemCtrl item_controller.ItemController) {
	flashareCtrl = FlashareController{
		authCtrl,
		itemCtrl,
	}
}
