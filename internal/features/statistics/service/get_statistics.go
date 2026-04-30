package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_errors "github.com/KKmanKK/golang-todoappTest/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userId *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"to must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statisticsRepository.GetTasks(
		ctx,
		userId,
		from,
		to,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"get task from repository:%w",
			err,
		)
	}

	statistics := calcStatistics(tasks)

	return statistics, nil
}

func calcStatistics(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.NewStatistics(0, 0, nil, nil)
	}

	tasksCreated := len(tasks)

	tasksComleted := 0
	var totalCompletionDuration time.Duration

	for _, task := range tasks {
		if task.Completed {
			tasksComleted++
		}
		completionDuration := task.CompletionDuration()
		if completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}

	tasksComletedRate := float64(tasksComleted) / float64(tasksCreated) * 100

	var taskAverageCompletionTime *time.Duration
	if tasksComleted > 0 && totalCompletionDuration != 0 {
		avg := totalCompletionDuration / time.Duration(tasksComleted)

		taskAverageCompletionTime = &avg
	}

	return domain.NewStatistics(
		tasksCreated,
		tasksComleted,
		&tasksComletedRate,
		taskAverageCompletionTime,
	)
}
