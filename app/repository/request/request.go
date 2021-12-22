package request_repository

import "flashare/entity"

type RequestRepository interface {
	GetPendingRequest(id string) ([]entity.Request, error)
	GetArchievedRequest(id string) ([]entity.Request, error)
}
