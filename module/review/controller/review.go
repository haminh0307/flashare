package review_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	review_controller "flashare/app/controller/review"
	review_usecase "flashare/app/usecase/review"
	user_usecase "flashare/app/usecase/user"
	"flashare/entity"
	flashare_errors "flashare/errors"
	"flashare/utils"
)

type reviewHandler struct {
	ReviewUC  review_usecase.ReviewUsecase
	ProfileUC user_usecase.ProfileUsecase
}

func NewReviewController(
	reviewUC review_usecase.ReviewUsecase,
	profileUC user_usecase.ProfileUsecase) review_controller.ReviewController {
	return &reviewHandler{
		reviewUC,
		profileUC,
	}
}

func (rvHandler *reviewHandler) SetupRouter(r *gin.RouterGroup) {
	r.POST("/add-review", rvHandler.AddReview)
	r.GET("/get-reviews", rvHandler.GetReviews)
}

type addReviewBody struct {
	Sender   string `json:"sender" binding:"required"`
	Receiver string `json:"receiver" binding:"required"`
	Rate     int    `json:"rate" binding:"required"`
	Review   string `json:"review"`
}

func (rvHandler *reviewHandler) AddReview(ctx *gin.Context) {
	var rv addReviewBody

	if err := ctx.ShouldBind(&rv); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	review := entity.Review{
		Sender:   rv.Sender,
		Receiver: rv.Review,
		Rate:     rv.Rate,
		Review:   rv.Review,
	}

	reviewId, err := rvHandler.ReviewUC.AddReview(review)

	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data:    reviewId,
	})
}

type reviewResponse struct {
	ReviewerId         string `json:"reviewer_id"`
	ReviewerName       string `json:"reviewer_name"`
	ReviewerAvatarLink string `json:"reviewer_avatar_link"`
	Rate               int    `json:"rate"`
	Review             string `json:"review"`
}

type profileReview struct {
	RateAvg float64          `json:"rate_avg"`
	Reviews []reviewResponse `json:"reviews"`
}

func (rvHandler *reviewHandler) GetReviews(ctx *gin.Context) {
	uid := ctx.Query("uid")
	if uid == "" {
		ctx.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    flashare_errors.ErrorInvalidParameters.Error(),
		})
		return
	}

	reviews, err := rvHandler.ReviewUC.GetReviews(uid)

	if err != nil {
		ctx.JSON(http.StatusOK, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
		return
	}

	rateAvg := 0.0
	reviewsResponse := []reviewResponse{}

	for _, rv := range reviews {
		rateAvg += float64(rv.Rate)

		user, err := rvHandler.ProfileUC.Get(rv.Sender)

		if err != nil {
			ctx.JSON(http.StatusOK, utils.DataResponse{
				Success: false,
				Data:    err.Error(),
			})
			return
		}

		review := reviewResponse{
			rv.Sender,
			user.FullName,
			user.AvatarLink,
			rv.Rate,
			rv.Review,
		}

		reviewsResponse = append(reviewsResponse, review)
	}

	if len(reviews) > 0 {
		rateAvg /= float64(len(reviews))
	}

	ctx.JSON(http.StatusOK, utils.DataResponse{
		Success: true,
		Data: profileReview{
			rateAvg,
			reviewsResponse,
		},
	})
}
