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

func (s *userService) UserInfo(ctx context.Context, login string) (*domain.User, error) {

	slog.Info("Service started \"UserInfo\"")

	user, err := s.repo.UserInfo(ctx, login)
	if err != nil {
		return nil, err
	}

	slog.Info("Service ended \"UserInfo\" success")
	return user, nil

}

func (s *userService) RegisterAccount(ctx context.Context, firstName, lastName, login, password string) (*domain.User, error) {

	slog.Info("Service started \"RegisterAccount\"")

	user, err := domain.NewUser(firstName, lastName, login, password)
	if err != nil {
		slog.Error("Service \"RegisterAccount\" get next error when create new user:", err)
		return nil, err
	}

	if err := s.repo.RegisterAccount(ctx, user); err != nil {
		return nil, err
	}

	slog.Info("Service ended \"RegisterAccount\" success")
	return user, nil

}
