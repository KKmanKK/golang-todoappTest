package task_transport_http

import (
	"time"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	Created_at   time.Time  `json:"created_at"`
	Completed_at *time.Time `json:"completed_at"`
	AuthorUserId int        `json:"author_user_id"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		Created_at:   task.Created_at,
		Completed_at: task.Completed_at,
		AuthorUserId: task.AuthorUserId,
	}
}

func tasksDTOFromDomains(tasks ...domain.Task) []TaskDTOResponse {
	tasksDTOResponse := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		tasksDTOResponse[i] = taskDTOFromDomain(task)
	}
	return tasksDTOResponse
}
