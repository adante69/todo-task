package taskManager

import "log/slog"

type Task struct {
	log             *slog.Logger
	taskSaver       TaskSaver
	taskProvider    TaskProvider
	commentSaver    CommentSaver
	commentProvider CommentProvider
	taskRemover     TaskRemover
}
