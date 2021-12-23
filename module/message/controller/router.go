package message_controller

import (
	"flashare/app/controller/message"
	"github.com/gin-gonic/gin"
)

type messageModule struct {
	MsgCtrl message_controller.MessageController
}

func NewMessageModule(msgCtrl message_controller.MessageController) message_controller.MessageModule {
	return &messageModule{
		msgCtrl,
	}
}

func (mMod *messageModule) SetupRouter(r *gin.RouterGroup) {
	messageRouter := r.Group("/message")

	mMod.MsgCtrl.SetupRouter(messageRouter)
}
