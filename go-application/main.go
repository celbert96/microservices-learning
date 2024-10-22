package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microservicesLearning/routes"
)

func main() {
	router := gin.Default()

	publicRoutes := router.Group("/public")
	routes.AddDefaultRoutes(publicRoutes)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("An error occurred when attempting to run the application > ", err)
		return
	}
}
