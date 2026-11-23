package storage_postgres

import (
	"context"
	"database/sql"
	"errors"

	"task-manager/internal/storage"
	"task-manager/internal/task"
)

type PostgresTaskStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresTaskStorage {
	return &PostgresTaskStorage{
		db: db,
	}
}

func (s *PostgresTaskStorage) Create(ctx context.Context, task *task.Task) error {
	query := `
        INSERT INTO tasks (title, description, completed)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `

	return s.db.QueryRowContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
	).Scan(&task.ID, &task.CreatedAt)
}

func (s *PostgresTaskStorage) GetByID(ctx context.Context, id int64) (*task.Task, error) {
	query := `
        SELECT id, title, description, completed, created_at
        FROM tasks
        WHERE id = $1
    `

	var task task.Task

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *PostgresTaskStorage) GetAll(ctx context.Context) ([]task.Task, error) {
	query := `
        SELECT id, title, description, completed, created_at
        FROM tasks
        ORDER BY id
    `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]task.Task, 0)

	for rows.Next() {
		var task task.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *PostgresTaskStorage) Update(ctx context.Context, task *task.Task) error {
	query := `
        UPDATE tasks
        SET title = $1,
            description = $2,
            completed = $3
        WHERE id = $4
    `

	result, err := s.db.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.ID,
	)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *PostgresTaskStorage) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Проверка соответствия интерфейсу на этапе компиляции.
var _ storage.TaskStorage = (*PostgresTaskStorage)(nil)
