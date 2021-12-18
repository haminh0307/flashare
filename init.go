package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"flashare/app/controller"
	"flashare/app/repository"
	"flashare/app/usecase"

	"flashare/module/user/controller"
	"flashare/module/user/repository"
	"flashare/module/user/usecase"

	"flashare/module/item/controller"
	"flashare/module/item/repository"
	"flashare/module/item/usecase"
)

func InitRepo(user, pwd, db string) {
	uri := "mongodb+srv://" + user + ":" + pwd + "@flashare.2hfwp.mongodb.net/" + db + "?retryWrites=true&w=majority"
	clientOptions := options.Client().
		ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user_repository.NewUserRepo(client.Database(db).Collection("users"))
	itemRepo := item_repository.NewItemRepo(client.Database(db).Collection("items"))

	repository.InitFlashareRepo(
		userRepo,
		itemRepo,
	)
}

func InitUsecase() {
	userRepo := repository.GetFlashareRepo().UserRepo
	authUC := user_usecase.NewAuthenticationUsecase(userRepo)

	itemRepo := repository.GetFlashareRepo().ItemRepo
	itemUC := item_usecase.NewItemUsecase(itemRepo)

	usecase.InitFlashareUsecase(
		authUC,
		itemUC,
	)
}

func InitController() {
	authUC := usecase.GetFlashareUsecase().AuthenticationUC
	authCtrl := user_controller.NewAuthenticationController(authUC)

	itemUC := usecase.GetFlashareUsecase().ItemUC
	itemCtrl := item_controller.NewItemController(itemUC)

	controller.InitFlashareController(
		authCtrl,
		itemCtrl,
	)
}

func Routing(r *gin.RouterGroup) {
	flashareController := controller.GetFlashareController()

	userModule := user_controller.NewUserModule(flashareController.AuthenticationCtrl)
	userModule.SetupRouter(r)

	itemModule := item_controller.NewItemModule(flashareController.ItemCtrl)
	itemModule.SetupRouter(r)
}
