package user_controller

import (
	"flashare/app/controller/user"
	"github.com/gin-gonic/gin"
)

type userModule struct {
	AuthenticationCtrl user_controller.AuthenticationController
}

func NewUserModule(authController user_controller.AuthenticationController) user_controller.UserModule {
	return &userModule{
		authController,
	}
}

func (uMod *userModule) SetupRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")

	uMod.AuthenticationCtrl.SetupRouter(userRouter.Group("/auth"))
}
