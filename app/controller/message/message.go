package message_controller

import (
	"github.com/gin-gonic/gin"
)

type MessageController interface {
	SetupRouter(r *gin.RouterGroup)
	FetchMessagesBetween(ctx *gin.Context)
}
