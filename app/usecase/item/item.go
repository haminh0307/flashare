package item_usecase

import (
	"flashare/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemUsecase interface {
	Fetch() ([]entity.Item, error)
	Upload(item entity.Item) (primitive.ObjectID, error)
}
