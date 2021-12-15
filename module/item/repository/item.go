package item_repository

import (
	"flashare/entity"
	"flashare/app/repository/item"
)

type itemRepoImpl struct {

}

func NewItemRepo() item_repository.ItemRepository {
	return &itemRepoImpl{}
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