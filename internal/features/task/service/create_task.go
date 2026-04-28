package task_service

import (
	"context"
	"fmt"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
)

func (s *TaskService) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{},
			fmt.Errorf(
				"validate task domain",
				err,
			)
	}

	task, err := s.taskRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{},
			fmt.Errorf(
				"create task error:%w",
				err,
			)
	}

	return task, nil
}
