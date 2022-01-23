package item_usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	item_repository "flashare/app/repository/item"
	item_usecase "flashare/app/usecase/item"
	"flashare/entity"
	flashare_errors "flashare/errors"
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

func (iUC *itemUsecaseImpl) FetchRandom(amount int) ([]entity.Item, error) {
	items, err := iUC.repo.FetchRandomOpenItem(amount)

	if err != nil {
		return nil, flashare_errors.ErrorFailToFetchRandomItem
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
