package user_controller

import (
	"github.com/gin-gonic/gin"
)

type AuthenticationController interface {
	SetupRouter(r *gin.RouterGroup)
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
}
