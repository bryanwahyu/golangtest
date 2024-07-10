package service

import (
  "errors"
  "github.com/bryanwahyu/test-golang/internal/domain"
  "github.com/bryanwahyu/test-golang/internal/repository"
)

type UserService struct {
  Repo *repository.UserRepository
}

func (s *UserService) Login(username, password string) (*domain.User, error) {
  user, err := s.Repo.ValidateUser(username, password)
  if err != nil {
    return nil, errors.New("invalid credentials")
  }
  return user, nil
}

func (s *UserService) CreateUser(user *domain.User) (*domain.User, error) {
  return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int32) (*domain.User, error) {
  return s.Repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *domain.User) (*domain.User, error) {
  return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int32) error {
  return s.Repo.DeleteUser(id)
}
