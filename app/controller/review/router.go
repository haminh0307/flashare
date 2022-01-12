package review_controller

import (
	"github.com/gin-gonic/gin"
)

type ReviewModule interface {
	SetupRouter(r *gin.RouterGroup)
}
