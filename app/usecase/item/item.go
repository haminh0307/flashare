package item_usecase

import (
	"flashare/entity"
)

type ItemUsecase interface {
	Fetch() ([]entity.Item, error)
}