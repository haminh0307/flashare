package request_usecase

import (
	"flashare/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestUsecase interface {
	GetPendingRequest(id string) (rqs []entity.Request, err error)
	GetArchievedRequest(id string) (rqs []entity.Request, err error)
	GetCancelledRequest(id string) (rqs []entity.Request, err error)
	SendRequest(userID string, itemID string) (entity.Request, error)
	GetItemRequest(itemID string) (rqs []entity.Request, err error)
	AcceptRequest(requestID primitive.ObjectID) (int64, error)
	CancelRequest(requestID primitive.ObjectID) (int64, error)
	ArchieveItem(itemID primitive.ObjectID) error
}
