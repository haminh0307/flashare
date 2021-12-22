package item_usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"flashare/app/repository/item"
	"flashare/app/usecase/item"
	"flashare/entity"
	"flashare/errors"
)

type itemUsecaseImpl struct {
	repo item_repository.ItemRepository
}

func NewItemUsecase(itemRepo item_repository.ItemRepository) item_usecase.ItemUsecase {
	return &itemUsecaseImpl{
		itemRepo,
	}
}

func (iUC *itemUsecaseImpl) Fetch(cate string) ([]entity.Item, error) {
	var items []entity.Item
	var err error

	// filter by category or not
	if cate == "" {
		items, err = iUC.repo.FetchOpenItem()
	} else {
		items, err = iUC.repo.FetchOpenItemByCategory(cate)
	}

	if err != nil {
		return nil, flashare_errors.ErrorFailToFetchItem
	}

	return items, err
}

func (iUC *itemUsecaseImpl) Upload(item entity.Item) (primitive.ObjectID, error) {
	item_id, err := iUC.repo.Create(item)
	if err != nil {
		return primitive.ObjectID{}, flashare_errors.ErrorFailToUploadItem
	}
	return item_id.(primitive.ObjectID), err
}
