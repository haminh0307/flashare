package item_controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	item_controller "flashare/app/controller/item"
	item_usecase "flashare/app/usecase/item"
	"flashare/entity"
	flashare_errors "flashare/errors"
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
	r.GET("/fetch-random", iHandler.FetchRandom)
	r.POST("/upload", iHandler.Upload)
}

func (iHandler *itemHandler) Fetch(ctx *gin.Context) {
	// use ctx.Query to match /fetch (all category) and /fetch?category=cate
	cate := ctx.Query("category")
	items, err := iHandler.ItemUC.Fetch(cate)
	if err != nil {
		// TODO: output?
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    items,
	})
}

func (iHandler *itemHandler) FetchRandom(ctx *gin.Context) {
	// use ctx.Query to match /fetch (all category) and /fetch?category=cate
	amt := ctx.Query("amount")

	amount, err := strconv.Atoi(amt);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	items, err := iHandler.ItemUC.FetchRandom(amount)

	if err != nil {
		// TODO: output?
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
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
	Title        string     `json:"title" binding:"required"`
	Category     string     `json:"category" binding:"required"`
	PhotosBase64 []string   `json:"photos_base64"`
	Description  string     `json:"description"`
	DueDate      *time.Time `json:"due_date,omitempty"`
	UploadedBy   string     `json:"uploaded_by" binding:"required"`
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

	photos_link := []string{}
	for _, photo := range rq.PhotosBase64 {
		link, err := utils.UploadBase64Image(photo)
		if err == nil {
			photos_link = append(photos_link, link)
		}
	}

	item := entity.Item{
		Title:       rq.Title,
		Category:    rq.Category,
		PhotosLink:  photos_link,
		Description: rq.Description,
		DueDate:     rq.DueDate,
		UploadedBy:  rq.UploadedBy,
		Status:      "open",
	}

	itemID, err := iHandler.ItemUC.Upload(item)

	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    itemID,
	})
}
