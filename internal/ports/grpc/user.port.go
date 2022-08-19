package portsGRPC

import (
	"context"
	"log"
	adapterPostgresRepo "togo/internal/adapter/postgresql/repositories"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"
	pb "togo/pkg/grpc"

	"github.com/go-pg/pg/v9"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SUserPort struct {
	UserService interface {
		CreateUser(user *domain.SUser) (*domain.SUser, error)
		GetAllUsers() ([]*domain.SUser, int, error)
		GetSingleUser(user *domain.SUser) error
		EditUser(user *domain.SUser) error
		DeleteUser(user *domain.SUser) error
	}
}

func NewUserPort(db *pg.DB) *SUserPort {
	userService := &serviceUsers.UserService{
		DB: db,
		UserRepo: &adapterPostgresRepo.SUserRepo{
			DB: db,
		},
	}

	return &SUserPort{
		UserService: userService,
	}
}

func (s *SUserPort) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	user := &domain.SUser{Username: in.GetUsername(), LimitPerDay: int(in.GetLimitPerDay())}

	newUser, err := s.UserService.CreateUser(user)

	pbUser := new(pb.User)
	pbUser.Id = newUser.ID
	pbUser.Username = newUser.Username
	pbUser.LimitPerDay = int64(newUser.LimitPerDay)
	pbUser.CreatedAt = timestamppb.New(newUser.CreatedAt)
	pbUser.UpdatedAt = timestamppb.New(newUser.UpdatedAt)

	log.Printf("Received: %v", pbUser)
	return &pb.CreateUserRes{User: pbUser}, err
}

func (s *SUserPort) GetAllUsers(ctx context.Context, in *pb.GetAllUsersReq) (*pb.GetAllUsersRes, error) {

	getUsers, total_user, err := s.UserService.GetAllUsers()

	users := make([]*pb.User, 0, len(getUsers))
	for _, user := range getUsers {
		newUser := new(pb.User)

		newUser.Id = user.ID
		newUser.Username = user.Username
		newUser.LimitPerDay = int64(user.LimitPerDay)
		newUser.CreatedAt = timestamppb.New(user.CreatedAt)
		newUser.UpdatedAt = timestamppb.New(user.UpdatedAt)

		users = append(users, newUser)
	}

	log.Printf("Received: %v", users)
	return &pb.GetAllUsersRes{User: users, TotalUser: int64(total_user)}, err
}

func (s *SUserPort) GetSingleUser(ctx context.Context, in *pb.GetSingleUserReq) (*pb.GetSingleUserRes, error) {
	user := &domain.SUser{ID: in.GetUserId()}
	err := s.UserService.GetSingleUser(user)

	newUser := new(pb.User)
	newUser.Id = user.ID
	newUser.Username = user.Username
	newUser.LimitPerDay = int64(user.LimitPerDay)
	newUser.CreatedAt = timestamppb.New(user.CreatedAt)
	newUser.UpdatedAt = timestamppb.New(user.UpdatedAt)

	log.Printf("Received: %v", newUser)
	return &pb.GetSingleUserRes{User: newUser}, err
}

func (s *SUserPort) EditUser(ctx context.Context, in *pb.EditUserReq) (*pb.EditUserRes, error) {
	user := &domain.SUser{ID: in.GetUserId()}

	err := s.UserService.EditUser(user)

	newUser := new(pb.User)
	newUser.Id = user.ID
	newUser.Username = user.Username
	newUser.LimitPerDay = int64(user.LimitPerDay)
	newUser.CreatedAt = timestamppb.New(user.CreatedAt)
	newUser.UpdatedAt = timestamppb.New(user.UpdatedAt)

	log.Printf("Received: %v", newUser)
	return &pb.EditUserRes{User: newUser}, err
}

func (s *SUserPort) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	user := &domain.SUser{ID: in.GetUserId()}
	err := s.UserService.DeleteUser(user)

	newUser := new(pb.User)
	newUser.Id = user.ID
	newUser.Username = user.Username
	newUser.LimitPerDay = int64(user.LimitPerDay)
	newUser.CreatedAt = timestamppb.New(user.CreatedAt)
	newUser.UpdatedAt = timestamppb.New(user.UpdatedAt)

	log.Printf("Received: %v", newUser)
	return &pb.DeleteUserRes{User: newUser}, err
}
