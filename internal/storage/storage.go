package storage

import (
	"context"
	"task-manager/internal/task"
)

type TaskStorage interface {
	Create(ctx context.Context, task *task.Task) error
	GetByID(ctx context.Context, id int64) (*task.Task, error)
	GetAll(ctx context.Context) ([]task.Task, error)
	Update(ctx context.Context, task *task.Task) error
	Delete(ctx context.Context, id int64) error
}
