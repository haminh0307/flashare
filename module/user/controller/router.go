package user_controller

import (
	"flashare/app/controller/user"
	"github.com/gin-gonic/gin"
)

type userModule struct {
	AuthenticationCtrl user_controller.AuthenticationController
	ProfileCtrl        user_controller.ProfileController
}

func NewUserModule(authController user_controller.AuthenticationController,
	profileController user_controller.ProfileController) user_controller.UserModule {
	return &userModule{
		authController,
		profileController,
	}
}

func (uMod *userModule) SetupRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")

	uMod.AuthenticationCtrl.SetupRouter(userRouter.Group("/auth"))
	uMod.ProfileCtrl.SetupRouter(userRouter.Group("/profile"))
}
