package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-task-api/internal/routes"
	"go-gin-task-api/pkg/logger"
)

func main() {
	r := gin.Default()

    logger.Init() // инициализируем логгер
	logger.Log.Info("Logger Init")

	routes.RegisterTaskRoutes(r)
	r.Run(":8080")
}