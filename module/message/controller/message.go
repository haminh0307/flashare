package message_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"flashare/app/controller/message"
	"flashare/app/usecase/message"
	"flashare/errors"
	"flashare/utils"
)

type messageHandler struct {
	MsgUC message_usecase.MessageUsecase
}

func NewMessageController(messageUC message_usecase.MessageUsecase) message_controller.MessageController {
	return &messageHandler{
		messageUC,
	}
}

func (mHandler *messageHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/fetch-between", mHandler.FetchMessagesBetween)
}

func (mHandler *messageHandler) FetchMessagesBetween(ctx *gin.Context) {
	user1_id := ctx.Query("user1")
	user2_id := ctx.Query("user2")

	// need to check null value
	if user1_id == "" || user2_id == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	msgs, err := mHandler.MsgUC.FetchMessagesBetween(user1_id, user2_id)
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
		Data:    msgs,
	})
}
