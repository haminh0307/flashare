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
	request_controller "flashare/module/request/controller"
	request_repository "flashare/module/request/repository"
	request_usecase "flashare/module/request/usecase"

	user_controller "flashare/module/user/controller"
	user_repository "flashare/module/user/repository"
	user_usecase "flashare/module/user/usecase"

	item_controller "flashare/module/item/controller"
	item_repository "flashare/module/item/repository"
	item_usecase "flashare/module/item/usecase"
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
	requestRepo := request_repository.NewRequestRepo(client.Database(db).Collection("requests"))

	repository.InitFlashareRepo(
		userRepo,
		itemRepo,
		requestRepo,
	)
}

func InitUsecase() {
	userRepo := repository.GetFlashareRepo().UserRepo
	authUC := user_usecase.NewAuthenticationUsecase(userRepo)

	itemRepo := repository.GetFlashareRepo().ItemRepo
	itemUC := item_usecase.NewItemUsecase(itemRepo)

	requestRepo := repository.GetFlashareRepo().RequestRepo
	requestUC := request_usecase.NewRequestUsecase(requestRepo)

	usecase.InitFlashareUsecase(
		authUC,
		itemUC,
		requestUC,
	)
}

func InitController() {
	authUC := usecase.GetFlashareUsecase().AuthenticationUC
	authCtrl := user_controller.NewAuthenticationController(authUC)

	itemUC := usecase.GetFlashareUsecase().ItemUC
	itemCtrl := item_controller.NewItemController(itemUC)

	requestUC := usecase.GetFlashareUsecase().RequestUC
	requestCtrl := request_controller.NewRequestController(requestUC)

	controller.InitFlashareController(
		authCtrl,
		itemCtrl,
		requestCtrl,
	)
}

func Routing(r *gin.RouterGroup) {
	flashareController := controller.GetFlashareController()

	userModule := user_controller.NewUserModule(flashareController.AuthenticationCtrl)
	userModule.SetupRouter(r)

	itemModule := item_controller.NewItemModule(flashareController.ItemCtrl)
	itemModule.SetupRouter(r)

	requestModule := request_controller.NewRequestModule(flashareController.RequestCtrl)
	requestModule.SetupRouter(r)
}
