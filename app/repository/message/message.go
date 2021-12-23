package message_repository

import (
	"flashare/entity"
)

type MessageRepository interface {
	FetchMessagesFromTo(sender_id, receiver_id string) ([]entity.Message, error)
}
