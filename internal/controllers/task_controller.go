package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-task-api/internal/models"
	"go-gin-task-api/internal/services"
	"go-gin-task-api/pkg/logger"
)

func AddTask(c *gin.Context) {
	logger.Log.Info("contoller AddTask")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := services.AddTask(task)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func GetTask(c *gin.Context) {
	logger.Log.Info("contoller GetTask")
	id := c.Param("id")
	task, err := services.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	logger.Log.Info("contoller GetTasks")
	tasks := services.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func HomePage(c *gin.Context) {
	logger.Log.Info("contoller HomePage")
	c.String(http.StatusOK, "Home page Task Manager")
}
