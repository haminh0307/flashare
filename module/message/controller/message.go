package message_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	message_controller "flashare/app/controller/message"
	message_usecase "flashare/app/usecase/message"
	user_usecase "flashare/app/usecase/user"
	flashare_errors "flashare/errors"
	"flashare/utils"
)

type messageHandler struct {
	MsgUC     message_usecase.MessageUsecase
	ProfileUC user_usecase.ProfileUsecase
}

func NewMessageController(
	messageUC message_usecase.MessageUsecase,
	profileUC user_usecase.ProfileUsecase) message_controller.MessageController {
	return &messageHandler{
		messageUC,
		profileUC,
	}
}

func (mHandler *messageHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/fetch-between", mHandler.FetchMessagesBetween)
	r.GET("/get-contacts", mHandler.GetContacts)
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

type contactResponse struct {
	UserId      string    `json:"uid"`
	Name        string    `json:"name"`
	AvatarLink  string    `json:"avatar_link"`
	LastMessage string    `json:"last_message"`
	Time        time.Time `json:"time"`
}

func (mHandler *messageHandler) GetContacts(ctx *gin.Context) {
	uid := ctx.Query("uid")

	if uid == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	messages, err := mHandler.MsgUC.GetContacts(uid)

	if err != nil {
		// TODO: output?
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	var contacts []contactResponse

	for _, msg := range messages {
		userId := msg.Sender
		if userId == uid {
			userId = msg.Receiver
		}

		user, err := mHandler.ProfileUC.Get(userId)

		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}

		contacts = append(contacts, contactResponse{
			userId,
			user.FullName,
			user.AvatarLink,
			msg.Content,
			msg.Time,
		})
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    contacts,
	})
}
