package item_controller

import (
	"flashare/app/controller/item"
	"github.com/gin-gonic/gin"
)

type itemModule struct {
	ItemCtrl item_controller.ItemController
}

func NewItemModule(itemCtrl item_controller.ItemController) itemModule {
	return itemModule{
		itemCtrl,
	}
}

func (iMod *itemModule) SetupRouter(r *gin.RouterGroup) {
	itemRouter := r.Group("/item")

	iMod.ItemCtrl.SetupRouter(itemRouter)
}
