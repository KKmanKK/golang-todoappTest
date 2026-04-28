package task_service

import (
	"context"
	"fmt"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_errors "github.com/KKmanKK/golang-todoappTest/internal/core/errors"
)

func (s *TaskService) GetTasks(
	ctx context.Context,
	userId *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negativ: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negativ: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	users, err := s.taskRepository.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}
	return users, nil
}
