package request_usecase

import (
	repository "flashare/app/repository"
	request_usecase "flashare/app/usecase/request"
	"flashare/entity"
	flashare_errors "flashare/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type requestUsecaseImpl struct {
	repo repository.FlashareRepo
}

func NewRequestUsecase(repo repository.FlashareRepo) request_usecase.RequestUsecase {
	return &requestUsecaseImpl{
		repo,
	}
}

func (rqUC *requestUsecaseImpl) GetPendingRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.repo.RequestRepo.GetPendingRequest(userID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}

func (rqUC *requestUsecaseImpl) GetArchievedRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.repo.RequestRepo.GetArchievedRequest(userID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}

func (rqUC *requestUsecaseImpl) SendRequest(userID string, itemID string) (rq entity.Request, err error) {
	res, err := rqUC.repo.ItemRepo.GetItemByID(itemID)
	if err != nil {
		err = flashare_errors.ErrorFailToFindItem
		return
	}
	_, err = rqUC.repo.RequestRepo.FindRequestByUserIDAndItemID(userID, itemID)
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
	ObjectId, err := rqUC.repo.RequestRepo.CreateRequest(rq)
	rq.ID = ObjectId.(primitive.ObjectID)
	if err != nil {
		err = flashare_errors.ErrorFailToCreateRequest
		return
	}
	return
}
