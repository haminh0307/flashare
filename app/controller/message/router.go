package message_controller

import (
	"github.com/gin-gonic/gin"
)

type MessageModule interface {
	SetupRouter(r *gin.RouterGroup)
}
