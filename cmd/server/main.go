package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karma/karma-backend/pkg/api/routes"

	"github.com/karma/karma-backend/internal/config"
	"github.com/karma/karma-backend/pkg/api/middlewares"
)

func main() {

	// Load env variables.
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	router := gin.Default()

	// Apply MongoDB middleware to all routes in taskApiGroup
	router.Use(middlewares.MongoDBMiddleware())

	taskApiGroup := router.Group("/api")

	routes.SetupTaskRoutes(taskApiGroup)

	port := ":8080"

	log.Printf("Server listening on %s", port)

	errRunningRouter := router.Run(port)

	if errRunningRouter != nil {
		log.Fatalf("Router cannot be run %v", errRunningRouter)
		return
	}
}