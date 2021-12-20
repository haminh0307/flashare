package item_repository

import (
	"flashare/entity"
)

type ItemRepository interface {
	Fetch() ([]entity.Item, error)
	Create(item entity.Item) (interface{}, error)
}
