package repository

import (
  "database/sql"
  "github.com/bryanwahyu/test-golang/internal/domain"
)

type UserRepository struct {
  DB *sql.DB
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
  var id int32
  err := r.DB.QueryRow("INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id",
    user.Name, user.Password, user.Email).Scan(&id)
  if err != nil {
    return nil, err
  }
  user.ID = id
  return user, nil
}

func (r *UserRepository) GetUserByID(id int32) (*domain.User, error) {
  var user domain.User
  err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
    Scan(&user.ID, &user.Name, &user.Email)
  if err != nil {
    return nil, err
  }
  return &user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
  _, err := r.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3",
    user.Name, user.Email, user.ID)
  if err != nil {
    return nil, err
  }
  return user, nil
}

func (r *UserRepository) DeleteUser(id int32) error {
  _, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
  return err
}

func (r *UserRepository) ValidateUser(email, password string) (*domain.User, error) {
  var user domain.User
  err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE email = $1 AND password = $2",
    email, password).Scan(&user.ID, &user.Name, &user.Email)
  if err != nil {
    return nil, err
  }
  return &user, nil
}
