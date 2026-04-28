package task_postgres_repository

import core_postgres_pool "github.com/KKmanKK/golang-todoappTest/internal/core/repository/postgres/pool"

type TaskRepository struct {
	pool core_postgres_pool.Pool
}

func NewTaskRepository(
	pool core_postgres_pool.Pool,
) *TaskRepository {
	return &TaskRepository{
		pool: pool,
	}
}
