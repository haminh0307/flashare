package review_controller

import (
	"github.com/gin-gonic/gin"
)

type ReviewController interface {
	SetupRouter(r *gin.RouterGroup)
	AddReview(ctx *gin.Context)
	GetReviews(ctx *gin.Context)
}
