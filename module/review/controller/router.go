package review_controller

import (
	review_controller "flashare/app/controller/review"

	"github.com/gin-gonic/gin"
)

type reviewModule struct {
	ReviewCtrl review_controller.ReviewController
}

func NewReviewModule(rvController review_controller.ReviewController) reviewModule {
	return reviewModule{
		rvController,
	}
}

func (rMod *reviewModule) SetupRouter(r *gin.RouterGroup) {
	rvRouter := r.Group("/review")

	rMod.ReviewCtrl.SetupRouter(rvRouter)
}
