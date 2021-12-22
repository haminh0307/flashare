package request_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	request_controller "flashare/app/controller/request"
	request_usecase "flashare/app/usecase/request"
	flashare_errors "flashare/errors"
	"flashare/utils"
)

type requestHandler struct {
	RequestUC request_usecase.RequestUsecase
}

func NewRequestController(requestUC request_usecase.RequestUsecase) request_controller.RequestController {
	return &requestHandler{
		requestUC,
	}
}

func (rqHandler *requestHandler) SetupRouter(r *gin.RouterGroup) {
	r.POST("/get-pending", rqHandler.GetPendingRequest)
	r.POST("/get-archieved", rqHandler.GetArchievedRequest)
}

type requestByUserID struct {
	UserID string `json:"user_id" binding:"required"`
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
	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    requestList,
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
	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    requestList,
	})
}
