package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"flashare/app/controller/user"
	"flashare/app/usecase/user"
	"flashare/errors"
	"flashare/utils"
)

type profileHandler struct {
	ProfileUC user_usecase.ProfileUsecase
}

func NewProfileController(authUC user_usecase.ProfileUsecase) user_controller.ProfileController {
	return &profileHandler{
		authUC,
	}
}

func (pHandler *profileHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("/get/:id", pHandler.Get)
	r.POST("/update-info", pHandler.UpdateInfo)
	r.POST("/update-avatar", pHandler.UpdateAvatar)
	r.POST("/change-password", pHandler.ChangePassword)
}

func (pHandler *profileHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := pHandler.ProfileUC.Get(id)
	// if error
	if err == flashare_errors.ErrorInternalServerError {
		ctx.JSON(http.StatusInternalServerError, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    user,
	})
}

type updateInfoRequest struct {
	ID          string `json:"id" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

func (pHandler *profileHandler) UpdateInfo(ctx *gin.Context) {
	var rq updateInfoRequest
	// if invalid parameters
	if err := ctx.ShouldBind(&rq); err != nil || rq.ID == "" || rq.FullName == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}
	if len(rq.PhoneNumber) > 0 { // if provided, must be valid
		valid, err := utils.IsValidPhoneNumber(rq.PhoneNumber)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.DataResponse{
				Success: false,
				Data:    flashare_errors.ErrorInternalServerError.Error(),
			})
			return
		}
		if !valid {
			ctx.JSON(http.StatusBadRequest, utils.DataResponse{
				Success: false,
				Data:    flashare_errors.ErrorInvalidParameters.Error(),
			})
			return
		}
	}

	err := pHandler.ProfileUC.UpdateInfo(rq.ID, rq.FullName, rq.PhoneNumber, rq.Address)
	// if server error
	if err == flashare_errors.ErrorInternalServerError {
		ctx.JSON(http.StatusInternalServerError, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInternalServerError.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
	})
}

type changePasswordRequest struct {
	ID     string `json:"id" binding:"required"`
	OldPwd string `json:"old_password" binding:"required"`
	NewPwd string `json:"new_password" binding:"required"`
}

func (pHandler *profileHandler) ChangePassword(ctx *gin.Context) {
	var rq changePasswordRequest
	// if invalid parameters
	if err := ctx.ShouldBind(&rq); err != nil || rq.ID == "" || rq.OldPwd == "" || rq.NewPwd == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	err := pHandler.ProfileUC.ChangePassword(rq.ID, rq.OldPwd, rq.NewPwd)
	// if server error
	if err == flashare_errors.ErrorInternalServerError {
		ctx.JSON(http.StatusInternalServerError, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInternalServerError.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
	})
}

type updateAvatarRequest struct {
	ID        string `json:"id" binding:"required"`
	AvtBase64 string `json:"avatar_base64" binding:"required"`
}

func (pHandler *profileHandler) UpdateAvatar(ctx *gin.Context) {
	var rq updateAvatarRequest
	// if invalid parameters
	if err := ctx.ShouldBind(&rq); err != nil || rq.ID == "" || rq.AvtBase64 == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	link, err := utils.UploadBase64Image(rq.AvtBase64)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	err = pHandler.ProfileUC.UpdateAvatar(rq.ID, link)
	// if server error
	if err == flashare_errors.ErrorInternalServerError {
		ctx.JSON(http.StatusInternalServerError, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInternalServerError.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
	})
}
