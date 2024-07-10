package grpc

import (
	"context"

	"github.com/bryanwahyu/test-golang/api/user"
	"github.com/bryanwahyu/test-golang/internal/domain"
	"github.com/bryanwahyu/test-golang/internal/service"
)

type UserServer struct {
  user.UnimplementedUserServiceServer
  UserService *service.UserService
}

func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
  newUser := &domain.User{
    Name: req.Name,
    Password: req.Password,
    Email:    req.Email,
  }
  createdUser, err := s.UserService.CreateUser(newUser)
  if err != nil {
    return nil, err
  }
  return &user.UserResponse{Id: createdUser.ID, Name: createdUser.Name, Email: createdUser.Email}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
  user, err := s.UserService.GetUserByID(req.Id)
  if err != nil {
    return nil, err
  }
  return &user.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error) {
  updatedUser := &domain.User{
    ID:       req.Id,
    Name: req.Name,
    Email:    req.Email,
  }
  user, err := s.UserService.UpdateUser(updatedUser)
  if err != nil {
    return nil, err
  }
  return &user.DefaultUserResponse{status:true,Message:"User berhasil di update"}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
  err := s.UserService.DeleteUser(req.Id)
  if err != nil {
    return nil, err
  }
  return &user.DefaultUserResponse{Message: "User deleted successfully"}, nil
}
