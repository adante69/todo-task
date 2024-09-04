package app

import (
	"log/slog"
	grpcapp "todo-task/internal/app/grpc"
	"todo-task/internal/services/taskManager"
	"todo-task/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dsn string,
) *App {
	storage, err := postgres.New(dsn)
	if err != nil {
		panic(err)
	}
	tmService := taskManager.NewTaskManager(log, storage, storage, storage, storage)

	GrpcApp := grpcapp.New(log, tmService, grpcPort)

	return &App{
		GRPCServer: GrpcApp,
	}
}
