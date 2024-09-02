package taskManager

import (
	"context"
	tmsv1 "github.com/adante69/todo-protos/gen/go/tms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskManager interface {
	Create(
		ctx context.Context,
		name string,
		description string,
		deadline string,
		priority string,
	) (taskId uint64, err error)
	Process(
		ctx context.Context,
		taskId uint64,
		status bool,
	) (answer bool, err error)
	Get(
		ctx context.Context,
		taskId uint64,
	) (name, description, deadline, priority, comment string, err error)
	AddComment(
		ctx context.Context,
		taskId uint64,
		comment string,
	) (answer bool, err error)
	Delete(
		ctx context.Context,
		taskId uint64,
	) (answer bool, err error)
}

type serverAPI struct {
	tmsv1.UnimplementedTaskControlServer
	taskManager TaskManager
}

func Register(gRpc *grpc.Server, taskManager TaskManager) {
	tmsv1.RegisterTaskControlServer(gRpc, &serverAPI{taskManager: taskManager})
}

func (s *serverAPI) Create(
	ctx context.Context,
	req *tmsv1.CreateRequest) (*tmsv1.CreateResponse, error) {

	if err := validateCreation(req); err != nil {
		return nil, err
	}

	id, err := s.taskManager.Create(ctx, req.GetName(), req.GetDescription(), req.GetDeadline(), req.GetPriority())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to create task")
	}

	return &tmsv1.CreateResponse{
		TaskId: id,
	}, nil
}

func (s *serverAPI) Process(
	ctx context.Context,
	req *tmsv1.ProcessingRequest,
) (*tmsv1.ProcessingResponse, error) {

	if req.GetTaskId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	answer, err := s.taskManager.Process(ctx, req.TaskId, req.GetStage())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to process task")
	}

	return &tmsv1.ProcessingResponse{
		Answer: answer,
	}, nil
}

func (s *serverAPI) Get(
	ctx context.Context,
	req *tmsv1.GetRequest,
) (*tmsv1.GetResponse, error) {
	if req.GetTaskId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	name, description, deadline, priority, comment, err := s.taskManager.Get(ctx, req.TaskId)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to get task")
	}

	return &tmsv1.GetResponse{
		Name:        name,
		Description: description,
		Deadline:    deadline,
		Priority:    priority,
		Comment:     comment,
	}, nil
}

func (s *serverAPI) AddComment(
	ctx context.Context,
	req *tmsv1.AddCommentRequest,
) (*tmsv1.AddCommentResponse, error) {
	if req.GetTaskId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	answer, err := s.taskManager.AddComment(ctx, req.GetTaskId(), req.GetComment())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to add comment")
	}

	return &tmsv1.AddCommentResponse{
		Answer: answer,
	}, nil
}

func (s *serverAPI) Delete(
	ctx context.Context,
	req *tmsv1.DeleteRequest,
) (*tmsv1.DeleteResponse, error) {
	if req.GetTaskId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	answer, err := s.taskManager.Delete(ctx, req.GetTaskId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to delete task")
	}

	return &tmsv1.DeleteResponse{
		Answer: answer,
	}, nil
}

func validateCreation(req *tmsv1.CreateRequest) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "empty name")
	}
	return nil
}
