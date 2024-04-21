package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karma/karma-backend/pkg/api/handlers"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	taskGroup := router.Group("/task")
	{
		taskGroup.GET("/get_task", handlers.GetTasks)
		taskGroup.POST("/create_task", handlers.CreateTask)
		taskGroup.PATCH("/update_task", handlers.UpdateTask)
		taskGroup.DELETE("/delete_task", handlers.DeleteTask)
		taskGroup.POST("/create_many", handlers.CreateMultipleTasks)
		// Add other user-related routes as needed
	}
}
