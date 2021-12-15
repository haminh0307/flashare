package item_usecase

import (
	"flashare/entity"
	"flashare/app/repository/item"
	"flashare/app/usecase/item"
)

type itemUsecaseImpl struct {
	repo item_repository.ItemRepository
}

func NewItemUsecase(itemRepo item_repository.ItemRepository) item_usecase.ItemUsecase {
	return &itemUsecaseImpl{
		itemRepo,
	}
}

func (iUC *itemUsecaseImpl) Fetch() ([]entity.Item, error) {
	items, err := iUC.repo.Fetch()
	if err != nil {
		return nil, ErrorFailToFetchItem
	}
	return items, err
}