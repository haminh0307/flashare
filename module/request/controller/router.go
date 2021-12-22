package request_controller

import (
	request_controller "flashare/app/controller/request"

	"github.com/gin-gonic/gin"
)

type requestModule struct {
	RequestCtrl request_controller.RequestController
}

func NewRequestModule(rqController request_controller.RequestController) requestModule {
	return requestModule{
		rqController,
	}
}

func (rMod *requestModule) SetupRouter(r *gin.RouterGroup) {
	rqRouter := r.Group("/request")

	rMod.RequestCtrl.SetupRouter(rqRouter)
}
