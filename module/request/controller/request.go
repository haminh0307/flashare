package request_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	request_controller "flashare/app/controller/request"
	item_usecase "flashare/app/usecase/item"
	request_usecase "flashare/app/usecase/request"
	user_usecase "flashare/app/usecase/user"
	"flashare/entity"
	flashare_errors "flashare/errors"
	"flashare/utils"
)

type requestHandler struct {
	RequestUC request_usecase.RequestUsecase
	ItemUC item_usecase.ItemUsecase
	ProfileUC user_usecase.ProfileUsecase
}

func NewRequestController(
	requestUC request_usecase.RequestUsecase,
	itemUC item_usecase.ItemUsecase,
	profileUC user_usecase.ProfileUsecase) request_controller.RequestController {
	return &requestHandler{
		requestUC,
		itemUC,
		profileUC,
	}
}

func (rqHandler *requestHandler) SetupRouter(r *gin.RouterGroup) {
	r.POST("/get-pending", rqHandler.GetPendingRequest)
	r.POST("/get-archieved", rqHandler.GetArchievedRequest)
	r.POST("/send-request", rqHandler.SendRequest)
	r.POST("/get-item-request", rqHandler.GetItemRequest)
}

type requestByUserID struct {
	UserID string `json:"user_id" binding:"required"`
}

type requestElement struct {
	Request entity.Request `json:"request"`
	Item entity.Item `json:"item"`
	Sender interface{} `json:"sender"`
}

func (rqHandler *requestHandler) GetPendingRequest(ctx *gin.Context) {
	var rq requestByUserID
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	requestList, err := rqHandler.RequestUC.GetPendingRequest(rq.UserID)
	// internal server error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	var data []requestElement

	for _, r := range requestList {
		
		item, err := rqHandler.ItemUC.GetItemById(r.Item)
		
		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}
		
		sender, err := rqHandler.ProfileUC.Get(r.Sender)

		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}

		data = append(data, requestElement{
			r,
			item,
			struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				AvatarLink string `json:"avatar_link"`
			}{
				r.Sender,
				sender.FullName,
				sender.AvatarLink,
			},
		})
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    data,
	})
}

func (rqHandler *requestHandler) GetArchievedRequest(ctx *gin.Context) {
	var rq requestByUserID
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	requestList, err := rqHandler.RequestUC.GetArchievedRequest(rq.UserID)
	// internal server error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}
	
	var data []requestElement

	for _, r := range requestList {
		
		item, err := rqHandler.ItemUC.GetItemById(r.Item)
		
		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}
		
		sender, err := rqHandler.ProfileUC.Get(r.Sender)

		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}

		data = append(data, requestElement{
			r,
			item,
			struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				AvatarLink string `json:"avatar_link"`
			}{
				r.Sender,
				sender.FullName,
				sender.AvatarLink,
			},
		})
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    data,
	})
}

type requestItem struct {
	UserID string `json:"user_id" binding:"required"`
	ItemID string `json:"item_id" binding:"required"`
}

func (rqHandler *requestHandler) SendRequest(ctx *gin.Context) {
	var rq requestItem
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	createdRequest, err := rqHandler.RequestUC.SendRequest(rq.UserID, rq.ItemID)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    createdRequest,
	})
}

type requestItemRequest struct {
	ItemID string `json:"item_id" binding:"required"`
}

func (rqHandler *requestHandler) GetItemRequest(ctx *gin.Context) {
	var rq requestItemRequest
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	requestList, err := rqHandler.RequestUC.GetItemRequest(rq.ItemID)
	// internal server error
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    requestList,
	})
}
