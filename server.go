package main

import (
	"log"
	"order-sys/repository"
	"order-sys/routes"
	"order-sys/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitRedis()

	service.StartKitchen(5)

	r := gin.Default()
	routes.SetRoutes(r)

	log.Println("Server is running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
