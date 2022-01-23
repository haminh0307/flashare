package request_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	request_repository "flashare/app/repository/request"
	"flashare/entity"
)

type requestRepoImpl struct {
	RequestColl *mongo.Collection
}

func NewRequestRepo(rcoll *mongo.Collection) request_repository.RequestRepository {
	return &requestRepoImpl{
		rcoll,
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

func (rRepo *requestRepoImpl) CreateRequest(request entity.Request) (interface{}, error) {
	rq, err := rRepo.RequestColl.InsertOne(context.Background(), request)
	return rq.InsertedID, err
}

func (rRepo *requestRepoImpl) FindRequestByUserIDAndItemID(userID, itemID string) (rq interface{}, err error) {
	filter := bson.D{{Key: "sender", Value: userID}, {Key: "item", Value: itemID}}
	err = rRepo.RequestColl.FindOne(context.Background(), filter).Decode(&rq)
	return
}

func (rRepo *requestRepoImpl) FindRequestByID(requestID primitive.ObjectID) (rq entity.Request, err error) {
	filter := bson.D{{Key: "_id", Value: requestID}}
	err = rRepo.RequestColl.FindOne(context.Background(), filter).Decode(&rq)
	return
}

func (rRepo *requestRepoImpl) GetItemRequest(itemID string) (rqs []entity.Request, err error) {
	filter := bson.D{{Key: "item", Value: itemID}}
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

func (rRepo *requestRepoImpl) AcceptRequest(requestID primitive.ObjectID) (int64, error) {
	filter := bson.D{{Key: "_id", Value: requestID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "accepted"}}}}
	res, err := rRepo.RequestColl.UpdateOne(context.Background(), filter, update)
	return res.ModifiedCount, err
}

func (rRepo *requestRepoImpl) CancelRequest(requestID primitive.ObjectID) (int64, error) {
	filter := bson.D{{Key: "_id", Value: requestID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "cancelled"}}}}
	res, err := rRepo.RequestColl.UpdateOne(context.Background(), filter, update)
	return res.ModifiedCount, err
}

func (rRepo *requestRepoImpl) CountAcceptedNumber(id string) (int64, error) {
	filter := bson.D{{Key: "item", Value: id}, {Key: "status", Value: "accepted"}}
	cnt, err := rRepo.RequestColl.CountDocuments(context.Background(), filter)
	return cnt, err
}

func (rRepo *requestRepoImpl) ArchieveRequest(id string) (int64, error) {
	filter := bson.D{{Key: "item", Value: id}, {Key: "status", Value: "accepted"}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "archieved"}}}}
	res, err := rRepo.RequestColl.UpdateOne(context.Background(), filter, update)
	return res.ModifiedCount, err
}
