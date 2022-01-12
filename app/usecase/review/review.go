package review_usecase

import (
	"flashare/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewUsecase interface {
	AddReview(entity.Review) (primitive.ObjectID, error)
	GetReviews(string) ([]entity.Review, error)
}
