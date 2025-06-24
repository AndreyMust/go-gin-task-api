package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-task-api/internal/controllers"
)

func RegisterTaskRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomePage)
	r.GET("/tasks/", controllers.GetTasks)
	r.POST("/task/", controllers.AddTask)
	r.GET("/task/:id", controllers.GetTask)
	r.DELETE("/task/:id", controllers.DeleteTask)
}
