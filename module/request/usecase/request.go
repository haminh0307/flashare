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

func (rqUC *requestUsecaseImpl) GetCancelledRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.rqRepo.GetCancelledRequest(userID)
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
	if res.Status != "open" {
		err = flashare_errors.ErrorRequestClosedItem
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

func (rqUC *requestUsecaseImpl) GetItemRequest(itemID string) ([]entity.Request, error) {
	rqs, err := rqUC.rqRepo.GetItemRequest(itemID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}

func (rqUC *requestUsecaseImpl) AcceptRequest(requestID primitive.ObjectID) (int64, error) {
	rq, err := rqUC.rqRepo.FindRequestByID(requestID)
	// fail to find request
	if err != nil {
		return 0, flashare_errors.ErrorFailToFindRequest
	}
	// can only accept a pending request
	if rq.Status != "pending" {
		return 0, flashare_errors.ErrorOnlyAcceptPendingRequest
	}
	cnt, err := rqUC.rqRepo.CountAcceptedNumber(rq.Item)
	// internal server error
	if err != nil {
		return 0, flashare_errors.ErrorInternalServerError
	}
	// only 1 request can be accepted
	if cnt >= 1 {
		return 0, flashare_errors.ErrorAcceptManyRequest
	}
	res, err := rqUC.rqRepo.AcceptRequest(requestID)
	// internal server error
	if err != nil {
		return 0, flashare_errors.ErrorInternalServerError
	}
	return res, err
}

func (rqUC *requestUsecaseImpl) CancelRequest(requestID primitive.ObjectID) (int64, error) {
	_, err := rqUC.rqRepo.FindRequestByID(requestID)
	// fail to find request
	if err != nil {
		return 0, flashare_errors.ErrorFailToFindRequest
	}
	res, err := rqUC.rqRepo.CancelRequest(requestID)
	// internal server error
	if err != nil {
		return 0, flashare_errors.ErrorInternalServerError
	}
	return res, err
}

func (rqUC *requestUsecaseImpl) ArchieveItem(itemID primitive.ObjectID) error {
	_, err := rqUC.itemRepo.GetItemByID(itemID)
	// fail to find item
	if err != nil {
		return flashare_errors.ErrorFailToFindItem
	}
	_, err = rqUC.rqRepo.ArchieveRequest(itemID.Hex())
	// internal server error
	if err != nil {
		return flashare_errors.ErrorInternalServerError
	}
	_, err = rqUC.itemRepo.ArchieveItem(itemID)
	// internal server error
	if err != nil {
		return flashare_errors.ErrorInternalServerError
	}
	return nil
}
