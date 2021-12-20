package item_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"flashare/app/repository/item"
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

func (iRepo *itemRepoImpl) Fetch() ([]entity.Item, error) {
	// TODO
	// SELECT * FROM Item
	// now fake
	return []entity.Item{
		{
			Title: "1",
		},
		{
			Title: "2",
		},
	}, nil
}

func (iRepo *itemRepoImpl) Create(item entity.Item) (interface{}, error) {
	res, err := iRepo.ItemColl.InsertOne(context.Background(), item)
	return res.InsertedID, err
}
