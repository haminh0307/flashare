package item_controller

import (
	"github.com/gin-gonic/gin"
)

type ItemModule interface {
	SetupRouter(r *gin.RouterGroup)
}
