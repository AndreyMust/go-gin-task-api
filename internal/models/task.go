package models

import "time"

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusRunning TaskStatus = "running"
	StatusDone    TaskStatus = "done"
	StatusFailed  TaskStatus = "failed"
)

type Task struct {
	ID      int `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	JobEndAt time.Time  `json:"job_ended_at"`
}