package request_usecase

import (
	request_repository "flashare/app/repository/request"
	request_usecase "flashare/app/usecase/request"
	"flashare/entity"
	flashare_errors "flashare/errors"
)

type requestUsecaseImpl struct {
	repo request_repository.RequestRepository
}

func NewRequestUsecase(rRepo request_repository.RequestRepository) request_usecase.RequestUsecase {
	return &requestUsecaseImpl{
		rRepo,
	}
}

func (rqUC *requestUsecaseImpl) GetPendingRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.repo.GetPendingRequest(userID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}

func (rqUC *requestUsecaseImpl) GetArchievedRequest(userID string) ([]entity.Request, error) {
	rqs, err := rqUC.repo.GetArchievedRequest(userID)
	// internal server error
	if err != nil {
		return nil, flashare_errors.ErrorInternalServerError
	}
	return rqs, err
}
