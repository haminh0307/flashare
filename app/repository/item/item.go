package item_repository

import (
	"flashare/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemRepository interface {
	FetchOpenItem() ([]entity.Item, error)
	FetchOpenItemByCategory(cate string) ([]entity.Item, error)
	FetchRandomOpenItem(amount int) ([]entity.Item, error)
	FetchItemUploadedBy(uid string) ([]entity.Item, error)
	Create(item entity.Item) (interface{}, error)
	GetItemByID(id primitive.ObjectID) (entity.Item, error)
	ArchieveItem(itemID primitive.ObjectID) (int64, error)
}
