package review_usecase

import (
	review_repository "flashare/app/repository/review"
	review_usecase "flashare/app/usecase/review"
	"flashare/entity"
	flashare_errors "flashare/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reviewUsecaseImpl struct {
	repo review_repository.ReviewRepository
}

func NewReviewUsecase(repo review_repository.ReviewRepository) review_usecase.ReviewUsecase {
	return &reviewUsecaseImpl{
		repo,
	}
}

func (rUC *reviewUsecaseImpl) AddReview(review entity.Review) (primitive.ObjectID, error) {
	reviewId, err := rUC.repo.Create(review)
	if err != nil {
		return primitive.ObjectID{}, flashare_errors.ErrorFailToAddReview
	}
	return reviewId.(primitive.ObjectID), err
}

func (rUC *reviewUsecaseImpl) GetReviews(userId string) ([]entity.Review, error) {
	reviews, err := rUC.repo.GetReviews(userId)
	if err != nil {
		return []entity.Review{}, flashare_errors.ErrorFailToGetReviews
	}
	return reviews, err
}
