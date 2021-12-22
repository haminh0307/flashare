package request_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	request_repository "flashare/app/repository/request"
	"flashare/entity"
)

type requestRepoImpl struct {
	RequestColl *mongo.Collection
}

func NewRequestRepo(ucoll *mongo.Collection) request_repository.RequestRepository {
	return &requestRepoImpl{
		ucoll,
	}
}

func (rRepo *requestRepoImpl) GetPendingRequest(userID string) (rqs []entity.Request, err error) {
	filter := bson.D{{Key: "sender", Value: userID}, {Key: "status", Value: bson.D{{Key: "$in", Value: bson.A{"pending", "accepted"}}}}}
	cursor, err := rRepo.RequestColl.Find(context.Background(), filter)
	if err != nil {
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var elem entity.Request
		err = cursor.Decode(&elem)
		if err != nil {
			return
		}
		rqs = append(rqs, elem)
	}
	return
}

func (rRepo *requestRepoImpl) GetArchievedRequest(userID string) (rqs []entity.Request, err error) {
	filter := bson.D{{Key: "sender", Value: userID}, {Key: "status", Value: "archieved"}}
	cursor, err := rRepo.RequestColl.Find(context.Background(), filter)
	if err != nil {
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var elem entity.Request
		err = cursor.Decode(&elem)
		if err != nil {
			return
		}
		rqs = append(rqs, elem)
	}
	return
}
