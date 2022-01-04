package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"flashare/app/repository/user"
	"flashare/entity"
	"flashare/utils"
)

type userRepoImpl struct {
	UserColl *mongo.Collection
}

func NewUserRepo(ucoll *mongo.Collection) user_repository.UserRepository {
	return &userRepoImpl{
		ucoll,
	}
}

func (uRepo *userRepoImpl) Get(id primitive.ObjectID) (u entity.User, err error) {
	filter := bson.D{{Key: "_id", Value: id}}
	err = uRepo.UserColl.FindOne(context.Background(), filter).Decode(&u)
	return
}

func (uRepo *userRepoImpl) GetByEmail(email string) (u entity.User, err error) {
	filter := bson.D{{Key: "email", Value: email}}
	err = uRepo.UserColl.FindOne(context.Background(), filter).Decode(&u)
	return
}

func (uRepo *userRepoImpl) Create(user entity.User) (interface{}, error) {
	res, err := uRepo.UserColl.InsertOne(context.Background(), user)
	return res.InsertedID, err
}

func (uRepo *userRepoImpl) Update(user entity.User) (bool, error) {
	// TODO
	filter := bson.D{{Key: "_id", Value: user.ID}}
	updateDoc, err := utils.ToDoc(user)
	if err != nil {
		return false, err
	}
	update := bson.D{{"$set", updateDoc}}
	res, err := uRepo.UserColl.UpdateOne(context.Background(), filter, update)
	return res.MatchedCount == 1, err
}
