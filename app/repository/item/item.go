package item_repository

import (
	"flashare/entity"
)

type ItemRepository interface {
	FetchOpenItem() ([]entity.Item, error)
	FetchOpenItemByCategory(cate string) ([]entity.Item, error)
	Create(item entity.Item) (interface{}, error)
	GetItemByID(id string) (entity.Item, error)
}
