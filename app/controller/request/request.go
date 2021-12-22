package request_controller

import (
	"github.com/gin-gonic/gin"
)

type RequestController interface {
	SetupRouter(r *gin.RouterGroup)
	GetPendingRequest(ctx *gin.Context)
	GetArchievedRequest(ctx *gin.Context)
}