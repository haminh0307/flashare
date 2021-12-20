package item_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"flashare/app/controller/item"
	"flashare/app/usecase/item"
	"flashare/utils"
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
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: true,
			Data:    err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    items,
	})
}
