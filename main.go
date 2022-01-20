package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	user_controller "flashare/module/user/controller"
	"flashare/utils"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	InitRepo(os.Getenv("MONGODB_USER"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_DATABASE"))
	InitUsecase()
	InitController()
	utils.InitCloudinary(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
}

func main() {
	router := gin.Default()
	Routing(router.Group("/api"))

	router.GET("/api/chat", user_controller.HandleChatConnection)
	router.GET("/api/chat/send", user_controller.HandleMessage)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
