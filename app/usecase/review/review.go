package review_usecase

import "flashare/entity"

type ReviewUsecase interface {
	AddReview(entity.Review) (interface{}, error)
	GetReviews(string) ([]entity.Review, error)
}
