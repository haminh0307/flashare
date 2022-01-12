package review_repository

import "flashare/entity"

type ReviewRepository interface {
	Create(entity.Review) (interface{}, error)
	GetReviews(string) ([]entity.Review, error)
}
