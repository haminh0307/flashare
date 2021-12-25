package request_usecase

import (
	item_repository "flashare/app/repository/item"
	request_repository "flashare/app/repository/request"
	request_usecase "flashare/app/usecase/request"
	"flashare/entity"
	flashare_errors "flashare/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type requestUsecaseImpl struct {
	rqRepo   request_repository.RequestRepository
	itemRepo item_repository.ItemRepository
}

func NewRequestUsecase(rqRepo request_repository.RequestRepository, itemRepo item_repository.ItemRepository) request_usecase.RequestUsecase {
	return &requestUsecaseImpl{
		rqRepo,
		itemRepo,
	}
}

func (rqUC *requestUsecaseImpl) GetPendingRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.rqRepo.GetPendingRequest(userID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}

func (rqUC *requestUsecaseImpl) GetArchievedRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.rqRepo.GetArchievedRequest(userID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}

func (rqUC *requestUsecaseImpl) SendRequest(userID string, itemID string) (rq entity.Request, err error) {
	objectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		err = flashare_errors.ErrorInvalidItemIdentity
		return
	}
	res, err := rqUC.itemRepo.GetItemByID(objectID)
	if err != nil {
		err = flashare_errors.ErrorFailToFindItem
		return
	}
	_, err = rqUC.rqRepo.FindRequestByUserIDAndItemID(userID, itemID)
	if err == nil {
		err = flashare_errors.ErrorRequestAlreadyExists
		return
	} else if err != mongo.ErrNoDocuments {
		err = flashare_errors.ErrorInternalServerError
		return
	}
	rq = entity.Request{
		Sender:   userID,
		Item:     itemID,
		Receiver: res.UploadedBy,
		Status:   "pending",
	}
	ObjectId, err := rqUC.rqRepo.CreateRequest(rq)
	rq.ID = ObjectId.(primitive.ObjectID)
	if err != nil {
		err = flashare_errors.ErrorFailToCreateRequest
		return
	}
	return
}
