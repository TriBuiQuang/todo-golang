package portsGRPC

import (
	"context"
	"fmt"
	"log"
	adapterPostgresRepo "togo/internal/adapter/postgresql/repositories"
	"togo/internal/core/domain"
	serviceTasks "togo/internal/core/services/tasks"
	serviceUsers "togo/internal/core/services/users"
	pb "togo/pkg/grpc"

	"github.com/go-pg/pg/v9"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type STaskPort struct {
	pb.UnimplementedTaskReaderServiceServer
	pb.UnimplementedTaskWriteServiceServer

	TaskService interface {
		CreateTask(task *domain.STask) (*domain.STask, error)
		GetAllTasks() ([]domain.STask, int, error)
		GetAllTaskToday(tasks *[]domain.STask, userId string) (*[]domain.STask, int, error)
		GetSingleTask(task *domain.STask) error
		DeleteTask(task *domain.STask) error
	}
}

func NewTaskPort(db *pg.DB) *STaskPort {
	taskService := &serviceTasks.TaskService{
		DB: db,
		TaskRepo: &adapterPostgresRepo.STaskRepo{
			DB: db,
		},
		UserService: serviceUsers.UserService{
			DB: db,
			UserRepo: &adapterPostgresRepo.SUserRepo{
				DB: db,
			},
		},
	}

	return &STaskPort{
		TaskService: taskService,
	}
}

func (s *STaskPort) CreateTask(ctx context.Context, in *pb.CreateTaskReq) (*pb.CreateTaskRes, error) {
	task := &domain.STask{UserID: in.GetUserId(), Name: in.GetName()}

	newTask, err := s.TaskService.CreateTask(task)

	pbTask := new(pb.Task)
	pbTask.Id = newTask.ID
	pbTask.UserId = newTask.UserID
	pbTask.Name = newTask.Name
	pbTask.CreatedAt = timestamppb.New(newTask.CreatedAt)
	pbTask.UpdatedAt = timestamppb.New(newTask.UpdatedAt)

	log.Printf("Received: %v", pbTask)
	return &pb.CreateTaskRes{Task: pbTask}, err
}

func (s *STaskPort) GetAllTasks(ctx context.Context, in *pb.GetAllTasksReq) (*pb.GetAllTasksRes, error) {
	fmt.Println("go here Get all tasks grpc")
	getTasks, total_task, err := s.TaskService.GetAllTasks()

	tasks := make([]*pb.Task, 0, len(getTasks))
	for _, task := range getTasks {
		newTask := new(pb.Task)

		newTask.Id = task.ID
		newTask.UserId = task.UserID
		newTask.Name = task.Name
		newTask.CreatedAt = timestamppb.New(task.CreatedAt)
		newTask.UpdatedAt = timestamppb.New(task.UpdatedAt)

		tasks = append(tasks, newTask)
	}

	log.Printf("Received: %v", tasks)
	return &pb.GetAllTasksRes{Task: tasks, TotalTask: int64(total_task)}, err
}

func (s *STaskPort) GetSingleTask(ctx context.Context, in *pb.GetSingleTaskReq) (*pb.GetSingleTaskRes, error) {
	task := &domain.STask{ID: in.GetTaskId()}
	err := s.TaskService.GetSingleTask(task)

	newTask := new(pb.Task)
	newTask.Id = task.ID
	newTask.UserId = task.UserID
	newTask.Name = task.Name
	newTask.CreatedAt = timestamppb.New(task.CreatedAt)
	newTask.UpdatedAt = timestamppb.New(task.UpdatedAt)

	log.Printf("Received: %v", newTask)
	return &pb.GetSingleTaskRes{Task: newTask}, err
}

func (s *STaskPort) DeleteTask(ctx context.Context, in *pb.DeleteTaskReq) (*pb.DeleteTaskRes, error) {
	task := &domain.STask{ID: in.GetTaskId()}
	err := s.TaskService.DeleteTask(task)

	newTask := new(pb.Task)
	newTask.Id = task.ID
	newTask.UserId = task.UserID
	newTask.Name = task.Name
	newTask.CreatedAt = timestamppb.New(task.CreatedAt)
	newTask.UpdatedAt = timestamppb.New(task.UpdatedAt)

	log.Printf("Received: %v", newTask)
	return &pb.DeleteTaskRes{Task: newTask}, err
}
