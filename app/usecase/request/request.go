package request_usecase

import "flashare/entity"

type RequestUsecase interface {
	GetPendingRequest(id string) (rqs []entity.Request, err error)
	GetArchievedRequest(id string) (rqs []entity.Request, err error)
}
