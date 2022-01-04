package user_controller

import (
	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	SetupRouter(r *gin.RouterGroup)
	Get(ctx *gin.Context)
	UpdateInfo(ctx *gin.Context)
	UpdateAvatar(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}
