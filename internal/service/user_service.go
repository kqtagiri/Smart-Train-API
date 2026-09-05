package service

import (
	"context"
	"log/slog"
	"smarttrain/internal/domain"
	"smarttrain/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {

	return &userService{
		repo: repo,
	}

}

func (s *userService) AllUsersInfo(ctx context.Context) (*[]domain.User, error) {

	slog.Info("Service started \"AllUsersInfo\"")

	users, err := s.repo.AllUsersInfo(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("Service ended \"AllUsersInfo\" success")
	return users, nil

}
