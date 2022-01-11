package message_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"flashare/app/repository/message"
	"flashare/entity"
)

type messageRepoImpl struct {
	MsgColl *mongo.Collection
}

func NewMessageRepo(coll *mongo.Collection) message_repository.MessageRepository {
	return &messageRepoImpl{
		coll,
	}
}

func (mRepo *messageRepoImpl) FetchMessagesFromTo(sender_id, receiver_id string) ([]entity.Message, error) {
	filter := bson.D{{Key: "sender", Value: sender_id}, {Key: "receiver", Value: receiver_id}}

	cursor, err := mRepo.MsgColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var res []entity.Message
	if err = cursor.All(context.Background(), &res); err != nil {
		return nil, err
	}

	return res, err
}

func (mRepo *messageRepoImpl) FetchMessages(uid string, is_sender bool) ([]entity.Message, error) {
	who := "sender"
	if !is_sender {
		who = "receiver"
	}

	filter := bson.D{{Key: who, Value: uid}}

	cursor, err := mRepo.MsgColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var res []entity.Message

	if err = cursor.All(context.Background(), &res); err != nil {
		return nil, err
	}

	return res, err
}
