package app

import (
	"log/slog"
	grpcapp "todo-task/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dsn string,
) *App {
	return &App{}
}
