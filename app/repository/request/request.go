package request_repository

import (
	"flashare/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestRepository interface {
	GetPendingRequest(id string) ([]entity.Request, error)
	GetArchievedRequest(id string) ([]entity.Request, error)
	GetCancelledRequest(id string) ([]entity.Request, error)
	CreateRequest(entity.Request) (interface{}, error)
	FindRequestByUserIDAndItemID(userID, itemID string) (interface{}, error)
	FindRequestByID(requestID primitive.ObjectID) (entity.Request, error)
	GetItemRequest(id string) ([]entity.Request, error)
	AcceptRequest(id primitive.ObjectID) (int64, error)
	CancelRequest(id primitive.ObjectID) (int64, error)
	CountAcceptedNumber(id string) (int64, error)
	ArchieveRequest(id string) (int64, error)
}
