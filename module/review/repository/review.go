package review_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	review_repository "flashare/app/repository/review"
	"flashare/entity"
)

type reviewRepoImpl struct {
	ReviewColl *mongo.Collection
}

func NewReviewRepo(rcoll *mongo.Collection) review_repository.ReviewRepository {
	return &reviewRepoImpl{
		rcoll,
	}
}

func (rRepo *reviewRepoImpl) Create(review entity.Review) (interface{}, error) {
	rv, err := rRepo.ReviewColl.InsertOne(context.Background(), review)
	return rv.InsertedID, err
}

func (rRepo *reviewRepoImpl) GetReviews(userId string) (reviews []entity.Review, err error) {
	filter := bson.D{{Key: "receiver", Value: userId}}
	cursor, err := rRepo.ReviewColl.Find(context.Background(), filter)
	if err != nil {
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var elem entity.Review
		err = cursor.Decode(&elem)
		if err != nil {
			return
		}
		reviews = append(reviews, elem)
	}
	return
}
