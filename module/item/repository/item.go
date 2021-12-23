package item_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	item_repository "flashare/app/repository/item"
	"flashare/entity"
)

type itemRepoImpl struct {
	ItemColl *mongo.Collection
}

func NewItemRepo(coll *mongo.Collection) item_repository.ItemRepository {
	return &itemRepoImpl{
		coll,
	}
}

func (iRepo *itemRepoImpl) FetchOpenItem() ([]entity.Item, error) {
	filter := bson.D{{Key: "status", Value: "open"}}

	cursor, err := iRepo.ItemColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var res []entity.Item
	if err = cursor.All(context.Background(), &res); err != nil {
		return nil, err
	}

	return res, err
}

func (iRepo *itemRepoImpl) FetchOpenItemByCategory(cate string) ([]entity.Item, error) {
	filter := bson.D{{Key: "status", Value: "open"}, {Key: "category", Value: cate}}

	cursor, err := iRepo.ItemColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var res []entity.Item
	if err = cursor.All(context.Background(), &res); err != nil {
		return nil, err
	}

	return res, err
}

func (iRepo *itemRepoImpl) Create(item entity.Item) (interface{}, error) {
	res, err := iRepo.ItemColl.InsertOne(context.Background(), item)
	return res.InsertedID, err
}

func (iRepo *itemRepoImpl) GetItemByID(id string) (res entity.Item, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	err = iRepo.ItemColl.FindOne(context.Background(), filter).Decode(&res)
	return
}
