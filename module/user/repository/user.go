package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"flashare/app/repository/user"
	"flashare/entity"
)

type userRepoImpl struct {
	UserColl *mongo.Collection
}

func NewUserRepo(ucoll *mongo.Collection) user_repository.UserRepository {
	return &userRepoImpl{
		ucoll,
	}
}

func (uRepo *userRepoImpl) GetByEmail(email string) (u entity.User, err error) {
	filter := bson.D{{"email", email}}
	err = uRepo.UserColl.FindOne(context.Background(), filter).Decode(&u)
	return
}

func (uRepo *userRepoImpl) Create(user entity.User) (interface{}, error) {
	res, err := uRepo.UserColl.InsertOne(context.Background(), user)
	return res.InsertedID, err
}
