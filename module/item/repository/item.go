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

func (iRepo *itemRepoImpl) FetchRandomOpenItem(amount int) ([]entity.Item, error) {
	filter := bson.D{{Key: "status", Value: "open"}}
	pipeline := []bson.M{bson.M{"$match": filter}, bson.M{"$sample": bson.M{"size": amount}}}

	cursor, err := iRepo.ItemColl.Aggregate(context.Background(), pipeline)

	if err != nil {
		return []entity.Item{}, err
	}

	var res []entity.Item

	for cursor.Next(context.Background()) {
		item := entity.Item{}
		err := cursor.Decode(&item)

		if err != nil {
			return []entity.Item{}, err
		} else {
			res = append(res, item)
		}
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

func (iRepo *itemRepoImpl) FetchItemUploadedBy(uid string) ([]entity.Item, error) {
	filter := bson.D{{Key: "uploaded_by", Value: uid}}

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

func (iRepo *itemRepoImpl) GetItemByID(id primitive.ObjectID) (res entity.Item, err error) {
	filter := bson.D{{Key: "_id", Value: id}}
	err = iRepo.ItemColl.FindOne(context.Background(), filter).Decode(&res)
	return
}
