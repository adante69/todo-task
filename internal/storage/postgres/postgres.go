package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"todo-task/internal/domain/models"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Task(ctx context.Context, taskID uint64) (task models.Task, err error) {
	const op = "storage.postgres.Task"

	query := `SELECT id, name, description, deadline, priority, end_status FROM tasks WHERE id = $1`

	row := s.db.QueryRowContext(ctx, query, taskID)

	err = row.Scan(&task.ID, &task.Name, &task.Desc, &task.Deadline, &task.Priority, &task.End_status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, fmt.Errorf("%s %w", op, err)
		}
		return models.Task{}, fmt.Errorf("%s %w", op, err)
	}
	return task, nil

}

func (s *Storage) CreateTask(ctx context.Context,
	name, description, deadline, priority string) (id uint64, err error) {
	const op = "storage.postgres.CreateTask"
	status := false

	query := `INSERT INTO tasks (name, description, deadline, priority, end_status) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = s.db.QueryRowContext(ctx, query, name, description, deadline, priority, status).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("%s %w", op, err)
	}
	return id, nil
}

func (s *Storage) Processing(ctx context.Context, id uint64) (bool, error) {
	const op = "storage.postgres.Process"

	query1 := `SELECT end_status FROM tasks WHERE id = $1`

	var status bool
	err := s.db.QueryRowContext(ctx, query1, id).Scan(&status)
	if err != nil {
		return false, fmt.Errorf("%s %w", op, err)
	}

	query := `
		UPDATE tasks
		SET end_status = $2
		WHERE id = $1
	`
	_, err = s.db.ExecContext(ctx, query, id, !status)
	if err != nil {
		return false, fmt.Errorf("%s %w", op, err)
	}

	return true, nil
}

func (s *Storage) UpdateTask(ctx context.Context, task models.Task) error {
	const op = "storage.postgres.UpdateTask"

	query := `
		UPDATE tasks
		SET name = $2, description = $3, deadline = $4, priority = $5
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, task.ID, task.Name, task.Desc, task.Deadline, task.Priority)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func New(source string) (*Storage, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	// Проверяем, что соединение действительно установлено и работает
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddComment(ctx context.Context, id uint64, comment string) (bool, error) {
	const op = "storage.postgres.AddComment"

	query := `INSERT INTO comments (id, comment) VALUES ($1, $2)`

	_, err := s.db.ExecContext(ctx, query, id, comment)
	if err != nil {
		return false, fmt.Errorf("%s %w", op, err)
	}
	return true, nil
}

func (s *Storage) RemoveComment(ctx context.Context, id uint64) error {
	const op = "storage.postgres.RemoveComment"
	query := `DELETE FROM comments WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}
	return nil
}

func (s *Storage) RemoveTask(ctx context.Context, id uint64) (bool, error) {
	const op = "storage.postgres.RemoveTask"
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("%s %w", op, err)
	}
	return true, nil
}
