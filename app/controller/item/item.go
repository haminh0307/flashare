package item_controller

import (
	"github.com/gin-gonic/gin"
)

type ItemController interface {
	SetupRouter(r *gin.RouterGroup)
	Fetch(ctx *gin.Context)
	Upload(ctx *gin.Context)
}
