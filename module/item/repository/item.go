package item_repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"flashare/app/repository/item"
	"flashare/entity"
)

type itemRepoImpl struct {
	Coll *mongo.Collection
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
			ID: "1",
		},
		{
			ID: "2",
		},
	}, nil
}
