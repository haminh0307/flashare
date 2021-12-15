package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	InitRepo()
	InitUsecase()
	InitController()
	router := gin.Default()
	Routing(router.Group("/api"))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}