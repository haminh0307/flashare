package message_usecase

import (
	"flashare/entity"
)

type MessageUsecase interface {
	FetchMessagesBetween(user1_id, user2_id string) ([]entity.Message, error)
	GetContacts(id string) ([]entity.Message, error)
}
