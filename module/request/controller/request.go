package request_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	ItemUC    item_usecase.ItemUsecase
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
	r.POST("/accept-request", rqHandler.AcceptRequest)
	r.POST("/cancel-request", rqHandler.CancelRequest)
	r.POST("/archieve-item", rqHandler.ArchieveItem)

}

type requestByUserID struct {
	UserID string `json:"user_id" binding:"required"`
}

type simpleUser struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AvatarLink string `json:"avatar_link"`
}

type requestElement struct {
	Request entity.Request `json:"request"`
	Item    entity.Item    `json:"item"`
	Sender  simpleUser     `json:"sender"`
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
			simpleUser{
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
			simpleUser{
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

type itemRequestElement struct {
	Request entity.Request `json:"request"`
	Sender  simpleUser     `json:"sender"`
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

	var data []itemRequestElement

	for _, r := range requestList {
		sender, err := rqHandler.ProfileUC.Get(r.Sender)

		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}

		data = append(data, itemRequestElement{
			r,
			simpleUser{
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

type modifyRequest struct {
	RequestID primitive.ObjectID `json:"request_id" binding:"required"`
}

func (rqHandler *requestHandler) AcceptRequest(ctx *gin.Context) {
	var rq modifyRequest
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	cnt, err := rqHandler.RequestUC.AcceptRequest(rq.RequestID)
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
		Data:    cnt,
	})
}

func (rqHandler *requestHandler) CancelRequest(ctx *gin.Context) {
	var rq modifyRequest
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	cnt, err := rqHandler.RequestUC.CancelRequest(rq.RequestID)
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
		Data:    cnt,
	})
}

type archieveItemRequest struct {
	ItemID primitive.ObjectID `json:"item_id" binding:"required"`
}

func (rqHandler *requestHandler) ArchieveItem(ctx *gin.Context) {
	var rq archieveItemRequest
	if err := ctx.ShouldBind(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	err := rqHandler.RequestUC.ArchieveItem(rq.ItemID)
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
		Data:    nil,
	})
}
