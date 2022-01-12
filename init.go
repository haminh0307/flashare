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

	message_controller "flashare/module/message/controller"
	message_repository "flashare/module/message/repository"
	message_usecase "flashare/module/message/usecase"

	review_controller "flashare/module/review/controller"
	review_repository "flashare/module/review/repository"
	review_usecase "flashare/module/review/usecase"
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
	messageRepo := message_repository.NewMessageRepo(client.Database(db).Collection("messages"))
	reviewRepo := review_repository.NewReviewRepo(client.Database(db).Collection("reviews"))

	repository.InitFlashareRepo(
		userRepo,
		itemRepo,
		requestRepo,
		messageRepo,
		reviewRepo,
	)
}

func InitUsecase() {
	userRepo := repository.GetFlashareRepo().UserRepo
	authUC := user_usecase.NewAuthenticationUsecase(userRepo)
	profileUC := user_usecase.NewProfileUsecase(userRepo)

	itemRepo := repository.GetFlashareRepo().ItemRepo
	itemUC := item_usecase.NewItemUsecase(itemRepo)

	requestRepo := repository.GetFlashareRepo().RequestRepo
	requestUC := request_usecase.NewRequestUsecase(requestRepo, itemRepo)

	messageRepo := repository.GetFlashareRepo().MessageRepo
	messageUC := message_usecase.NewMessageUsecase(messageRepo)

	reviewRepo := repository.GetFlashareRepo().ReviewRepo
	reviewUC := review_usecase.NewReviewUsecase(reviewRepo)

	usecase.InitFlashareUsecase(
		authUC,
		profileUC,
		itemUC,
		requestUC,
		messageUC,
		reviewUC,
	)
}

func InitController() {
	authUC := usecase.GetFlashareUsecase().AuthenticationUC
	authCtrl := user_controller.NewAuthenticationController(authUC)

	profileUC := usecase.GetFlashareUsecase().ProfileUC
	profileCtrl := user_controller.NewProfileController(profileUC)

	itemUC := usecase.GetFlashareUsecase().ItemUC
	itemCtrl := item_controller.NewItemController(itemUC)

	requestUC := usecase.GetFlashareUsecase().RequestUC
	requestCtrl := request_controller.NewRequestController(requestUC)

	messageUC := usecase.GetFlashareUsecase().MessageUC
	messageCtrl := message_controller.NewMessageController(messageUC, profileUC)

	reviewUC := usecase.GetFlashareUsecase().ReviewUC
	reviewCtrl := review_controller.NewReviewController(reviewUC, profileUC)

	controller.InitFlashareController(
		authCtrl,
		profileCtrl,
		itemCtrl,
		requestCtrl,
		messageCtrl,
		reviewCtrl,
	)
}

func Routing(r *gin.RouterGroup) {
	flashareController := controller.GetFlashareController()

	userModule := user_controller.NewUserModule(flashareController.AuthenticationCtrl, flashareController.ProfileCtrl)
	userModule.SetupRouter(r)

	itemModule := item_controller.NewItemModule(flashareController.ItemCtrl)
	itemModule.SetupRouter(r)

	requestModule := request_controller.NewRequestModule(flashareController.RequestCtrl)
	requestModule.SetupRouter(r)

	messageModule := message_controller.NewMessageModule(flashareController.MessageCtrl)
	messageModule.SetupRouter(r)

	reviewModule := review_controller.NewReviewModule(flashareController.ReviewCtrl)
	reviewModule.SetupRouter(r)
}
