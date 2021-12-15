package main

import (
	"github.com/gin-gonic/gin"

	"flashare/app/repository"
	"flashare/app/usecase"
	"flashare/app/controller"

	"flashare/module/item/repository"
	"flashare/module/item/usecase"
	"flashare/module/item/controller"
)

func InitRepo() {
	itemRepo := item_repository.NewItemRepo()

	repository.InitFlashareRepo(
		itemRepo,
	)
}

func InitUsecase() {
	itemRepo := repository.GetFlashareRepo().ItemRepo
	itemUC := item_usecase.NewItemUsecase(itemRepo)

	usecase.InitFlashareUsecase(
		itemUC,
	)
}

func InitController() {
	itemUC := usecase.GetFlashareUsecase().ItemUC
	itemController := item_controller.NewItemController(itemUC)

	controller.InitFlashareController(
		itemController,
	)
}

func Routing(r *gin.RouterGroup) {
	flashareController := controller.GetFlashareController()
	flashareController.ItemController.SetupRouter(r)
}