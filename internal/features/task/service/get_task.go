package task_service

import (
	"context"
	"fmt"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
)

func (s *TaskService) GetTask(
	ctx context.Context,
	taskId int,
) (domain.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, taskId)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"get task from repository: %w",
			err,
		)
	}
	return task, nil
}
