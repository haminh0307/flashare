package item_controller

import (
	"net/http"

	"flashare/app/controller/item"
	"flashare/app/usecase/item"
	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	ItemUC item_usecase.ItemUsecase
}

func NewItemController(itemUC item_usecase.ItemUsecase) item_controller.ItemController {
	return &itemHandler{
		itemUC,
	}
}

func (iHandler *itemHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/fetch", iHandler.Fetch)
}

func (iHandler *itemHandler) Fetch(ctx *gin.Context) {
	items, err := iHandler.ItemUC.Fetch()
	if err != nil {
		// TODO: output?
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "okay", "data": items})
}
