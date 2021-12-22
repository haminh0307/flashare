package user_controller

import (
	"github.com/gin-gonic/gin"
)

type UserModule interface {
	SetupRouter(r *gin.RouterGroup)
}
