package request_repository

import "flashare/entity"

type RequestRepository interface {
	GetPendingRequest(id string) ([]entity.Request, error)
	GetArchievedRequest(id string) ([]entity.Request, error)
	CreateRequest(entity.Request) (interface{}, error)
	FindRequestByUserIDAndItemID(userID, itemID string) (interface{}, error)
	GetItemRequest(id string) ([]entity.Request, error)
}
