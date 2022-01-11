package message_usecase

import (
	"sort"

	"flashare/app/repository/message"
	"flashare/app/usecase/message"
	"flashare/entity"
	"flashare/errors"
)

type messageUsecaseImpl struct {
	repo message_repository.MessageRepository
}

func NewMessageUsecase(msgRepo message_repository.MessageRepository) message_usecase.MessageUsecase {
	return &messageUsecaseImpl{
		msgRepo,
	}
}

func (mUC *messageUsecaseImpl) FetchMessagesBetween(user1_id, user2_id string) ([]entity.Message, error) {
	// from user1 to user2
	msg12, err := mUC.repo.FetchMessagesFromTo(user1_id, user2_id)

	if err != nil {
		return nil, flashare_errors.ErrorFailToFetchMessage
	}

	// from user2 to user1
	msg21, err := mUC.repo.FetchMessagesFromTo(user2_id, user1_id)

	if err != nil {
		return nil, flashare_errors.ErrorFailToFetchMessage
	}

	res := append(msg12, msg21...)

	sort.Slice(res, func(i, j int) bool {
		return res[i].Time.After(res[j].Time)
	})

	return res, err
}

func (mUC *messageUsecaseImpl) GetContacts(uid string) ([]entity.Message, error) {
	// from uid
	msg1, err := mUC.repo.FetchMessages(uid, true)

	if err != nil {
		return nil, flashare_errors.ErrorFailToGetContacts
	}

	// to uid
	msg2, err := mUC.repo.FetchMessages(uid, false)

	if err != nil {
		return nil, flashare_errors.ErrorFailToGetContacts
	}

	res := append(msg1, msg2...)

	sort.Slice(res, func(i, j int) bool {
		return res[i].Time.After(res[j].Time)
	})

	var messages []entity.Message
	set := map[string]bool{}

	for _, msg := range res {
		contact := msg.Sender
		if contact == uid {
			contact = msg.Receiver
		}
		if !set[contact] {
			set[contact] = true
			messages = append(messages, msg)
		}
	}

	return messages, err
}
