package services

import (
	"strconv"
	"time"
	"sync"
	"errors"

	"go-gin-task-api/internal/models"
	"go-gin-task-api/pkg/logger"
)

var tasks = make(map[int]models.Task)

var	mu    sync.RWMutex
var	task_cnt int

func AddTask(task models.Task) (models.Task, error) {
	logger.Log.Info("services AddTask")
	task.CreatedAt = time.Now()
	mu.Lock()
	defer mu.Unlock()

	task_cnt++
	task.ID = task_cnt

	tasks[task.ID] = task
	go runTask(&task)
	return task, nil
}

func DeleteTask(id string) error {
	logger.Log.Info("services DeleteTask")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[intID]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, intID)
	return nil
}

func GetTaskById(id string) (*models.Task, error) {
	logger.Log.Info("services GetTaskById")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	mu.RLock()
	defer mu.RUnlock()

	task, exists := tasks[intID]
	if !exists {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func GetTasks() []models.Task {
	logger.Log.Info("services GetTasks")
	mu.RLock()
	defer mu.RUnlock()

	result := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, t)
	}
	return result
}


func runTask(task *models.Task) {
	logger.Log.Info("services runTask")
	UpdateTaskStatus(task.ID, models.StatusRunning)
	time.Sleep(20 * time.Second)      // 20 секунды
	UpdateTaskStatus(task.ID, models.StatusDone)
}


func UpdateTaskStatus(id int, status models.TaskStatus) {
	logger.Log.Info("services runTask", "id", id, "status ", status)
	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[id]
	if exists {
		if status == models.StatusDone {
			task.JobEndAt = time.Now()
		}
		task.Status = status
		tasks[id] = task 
	}
}
