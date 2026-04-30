package statistics_postgres_repository

import (
	"time"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
	AuthorUserId int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.Created_at,
		taskModel.Completed_at,
		taskModel.AuthorUserId,
	)
}

func tasksDomaInFromModels(tasksModeles []TaskModel) []domain.Task {
	tasksDomains := make([]domain.Task, len(tasksModeles))

	for i, taskModel := range tasksModeles {
		tasksDomains[i] = taskDomainFromModel(taskModel)
	}

	return tasksDomains
}
