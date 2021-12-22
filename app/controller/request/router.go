package request_controller

import (
	"github.com/gin-gonic/gin"
)

type RequestModule interface {
	SetupRouter(r *gin.RouterGroup)
}
