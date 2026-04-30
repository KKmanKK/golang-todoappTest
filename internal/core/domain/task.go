package domain

import (
	"fmt"
	"time"

	core_errors "github.com/KKmanKK/golang-todoappTest/internal/core/errors"
)

type Task struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
	AuthorUserId int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	created_at time.Time,
	completed_at *time.Time,
	authorUserId int,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		Created_at:   created_at,
		Completed_at: completed_at,
		AuthorUserId: authorUserId,
	}
}

func NewTaskUninitialized(
	title string,
	description *string,
	authorUserId int,
) Task {
	return NewTask(
		UnitializedID,
		UnitializedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		authorUserId,
	)
}

func (t *Task) CompletionDuration() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.Completed_at == nil {
		return nil
	}

	duration := t.Completed_at.Sub(t.Created_at)

	return &duration
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"ivalid `Title` len: %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"ivalid `Description` len: %d: %w",
				descriptionLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}
	if t.Completed {
		if t.Completed_at == nil {
			return fmt.Errorf(
				"`CompletedAt` can't be 'nil' if 'Completed'=true: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		if t.Completed_at.Before(t.Created_at) {
			return fmt.Errorf(
				"'CompletedAt' can't be befor 'CreatedAt': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	} else {
		if t.Completed_at != nil {
			return fmt.Errorf(
				"`CompletedAt` must be 'nil' if 'Completed'=false: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("'Title' can't be patched to Null: %w", core_errors.ErrInvalidArgument)
	}
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("'Completed' can't be patched to Null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", err)
	}

	tmp := *u

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value

		if tmp.Completed {
			completedAt := time.Now()
			tmp.Completed_at = &completedAt
		} else {
			tmp.Completed_at = nil
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}

	*u = tmp

	return nil
}
