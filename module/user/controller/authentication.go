package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	user_controller "flashare/app/controller/user"
	user_usecase "flashare/app/usecase/user"
	"flashare/entity"
	flashare_errors "flashare/errors"
	"flashare/utils"
)

type authenticationHandler struct {
	AuthenticationUC user_usecase.AuthenticationUsecase
}

func NewAuthenticationController(authUC user_usecase.AuthenticationUsecase) user_controller.AuthenticationController {
	return &authenticationHandler{
		authUC,
	}
}

func (authHandler *authenticationHandler) SetupRouter(r *gin.RouterGroup) {
	r.POST("/sign-in", authHandler.SignIn)
	r.POST("/sign-up", authHandler.SignUp)
}

type signInRequest struct {
	Email string `json:"email" binding:"required"`
	Pwd   string `json:"password" binding:"required"`
}

func (authHandler *authenticationHandler) SignIn(ctx *gin.Context) {
	var rq signInRequest
	// if invalid parameters
	if err := ctx.ShouldBind(&rq); err != nil || rq.Email == "" || rq.Pwd == "" || !utils.IsValidMail(rq.Email) {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	user, err := authHandler.AuthenticationUC.SignIn(rq.Email, rq.Pwd)
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

type signUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Pwd      string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

func (authHandler *authenticationHandler) SignUp(ctx *gin.Context) {
	var rq signUpRequest
	// if invalid parameters
	if err := ctx.ShouldBind(&rq); err != nil || rq.Email == "" || rq.Pwd == "" || rq.FullName == "" || !utils.IsValidMail(rq.Email) {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	pwd_hash, err := bcrypt.GenerateFromPassword([]byte(rq.Pwd), bcrypt.MinCost)
	// if server error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInternalServerError.Error(),
		})
		return
	}

	u := entity.User{
		Email:            rq.Email,
		PasswordHashCode: pwd_hash,
		FullName:         rq.FullName,
	}

	user_id, err := authHandler.AuthenticationUC.SignUp(u)
	if err != nil {
		// TODO: output?
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    user_id,
	})
}
