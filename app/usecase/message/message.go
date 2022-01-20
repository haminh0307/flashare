package message_usecase

import (
	"flashare/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageUsecase interface {
	FetchMessagesBetween(user1_id, user2_id string) ([]entity.Message, error)
	GetContacts(id string) ([]entity.Message, error)
	AddMessage(entity.Message) (primitive.ObjectID, error)
}
