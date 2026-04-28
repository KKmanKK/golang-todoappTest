package task_service

import (
	"context"
	"fmt"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
)

func (s *TaskService) PatchTask(
	ctx context.Context,
	id int,
	patch domain.TaskPatch,
) (domain.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get tasks: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	task, err = s.taskRepository.PatchTask(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task:%w", err)
	}
	return task, nil
}
