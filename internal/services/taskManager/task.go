package taskManager

import (
	"context"
	"log/slog"
	"todo-task/internal/domain/models"
)

type TaskManager struct {
	log             *slog.Logger
	taskSaver       TaskSaver
	taskProvider    TaskProvider
	commentSaver    CommentSaver
	commentProvider CommentSaver
	taskRemover     TaskRemover
}

type TaskSaver interface {
	CreateTask(ctx context.Context, name, description, deadline, priority string) (uint64, error)
	UpdateTask(ctx context.Context, task models.Task) error
	Processing(ctx context.Context, id uint64) (bool, error)
}

type TaskProvider interface {
	Task(ctx context.Context, taskID uint64) (models.Task, error)
}

type CommentSaver interface {
	AddComment(ctx context.Context, id uint64, comment string) (bool, error)
	RemoveComment(ctx context.Context, id uint64) error
}

type TaskRemover interface {
	RemoveTask(ctx context.Context, id uint64) (bool, error)
}

func NewTaskManager(
	log *slog.Logger,
	taskSaver TaskSaver,
	taskProvider TaskProvider,
	commentSaver CommentSaver,
	taskRemover TaskRemover,
) *TaskManager {
	return &TaskManager{
		log:          log,
		taskSaver:    taskSaver,
		taskProvider: taskProvider,
		commentSaver: commentSaver,
		taskRemover:  taskRemover,
	}
}

func (tm *TaskManager) Create(ctx context.Context, name, description, deadline, priority string) (uint64, error) {
	tm.log.Info("Creating task")
	return tm.taskSaver.CreateTask(ctx, name, description, deadline, priority)
}

func (tm *TaskManager) Get(ctx context.Context, taskID uint64) (models.Task, error) {
	tm.log.Info("Fetching task")
	task, err := tm.taskProvider.Task(ctx, taskID)
	if err != nil {
		return models.Task{}, err
	}

	return task, err
}

func (tm *TaskManager) UpdateTask(ctx context.Context, task models.Task) error {
	tm.log.Info("Updating task")
	return tm.taskSaver.UpdateTask(ctx, task)
}

func (tm *TaskManager) AddComment(ctx context.Context, taskID uint64, comment string) (bool, error) {
	tm.log.Info("Adding comment")
	return tm.commentSaver.AddComment(ctx, taskID, comment)
}

func (tm *TaskManager) RemoveComment(ctx context.Context, taskID uint64) error {
	tm.log.Info("Removing comment")
	return tm.commentSaver.RemoveComment(ctx, taskID)
}

func (tm *TaskManager) Delete(ctx context.Context, taskID uint64) (bool, error) {
	tm.log.Info("Removing task")
	return tm.taskRemover.RemoveTask(ctx, taskID)
}

func (tm *TaskManager) Processing(ctx context.Context, id uint64) (bool, error) {
	tm.log.Info("Processing task")
	return tm.taskSaver.Processing(ctx, id)
}
