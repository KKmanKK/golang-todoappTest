package task_service

import (
	"context"
	"fmt"
)

func (s *TaskService) DeleteTask(
	ctx context.Context,
	taskId int,
) error {
	err := s.taskRepository.DeleteTask(ctx, taskId)
	if err != nil {
		return fmt.Errorf("delete task from repository: %w", err)
	}
	return nil
}
