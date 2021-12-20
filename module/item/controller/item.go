package item_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"flashare/app/controller/item"
	"flashare/app/usecase/item"
	"flashare/entity"
	"flashare/errors"
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
	r.POST("/upload", iHandler.Upload)
}

func (iHandler *itemHandler) Fetch(ctx *gin.Context) {
	items, err := iHandler.ItemUC.Fetch()
	if err != nil {
		// TODO: output?
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: true,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    items,
	})
}

type uploadRequest struct {
	Title       string     `json:"title"`
	Category    string     `json:"category"`
	PhotosLink  []string   `json:"photos_link"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	UploadedBy  string     `json:"uploaded_by"`
}

func (iHandler *itemHandler) Upload(ctx *gin.Context) {
	var rq uploadRequest
	// if invalid parameters
	if err := ctx.ShouldBind(&rq); err != nil || rq.Title == "" || rq.Category == "" || rq.UploadedBy == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	item_id, err := iHandler.ItemUC.Upload(entity.Item{
		Title:       rq.Title,
		Category:    rq.Category,
		PhotosLink:  rq.PhotosLink,
		Description: rq.Description,
		DueDate:     rq.DueDate,
		UploadedBy:  rq.UploadedBy,
		Status:      "open",
	})

	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: true,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    item_id,
	})
}
