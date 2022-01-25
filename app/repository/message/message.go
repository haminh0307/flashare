package message_repository

import (
	"flashare/entity"
)

type MessageRepository interface {
	FetchMessagesFromTo(sender_id, receiver_id string) ([]entity.Message, error)
	FetchMessages(uid string, is_sender bool) ([]entity.Message, error)
	CreateMessage(entity.Message) (interface{}, error)
}
