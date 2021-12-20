package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
