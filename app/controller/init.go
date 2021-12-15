package controller

import (
	"flashare/app/controller/item"
)

type FlashareController struct {
	ItemController item_controller.ItemController
}

var flashareController FlashareController

func GetFlashareController() FlashareController {
	return flashareController
}

func InitFlashareController(itemController item_controller.ItemController) {
	flashareController = FlashareController{
		itemController,
	}
}