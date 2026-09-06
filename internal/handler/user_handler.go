package handler

import (
	"context"
	"errors"
	"log/slog"
	"smarttrain/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	AllUsersInfo(ctx context.Context) (*[]domain.User, error)
	UserInfo(ctx context.Context, login string) (*domain.User, error)
	RegisterAccount(ctx context.Context, firstName, lastName, login, password string) (*domain.User, error)
}

type userHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *userHandler {

	return &userHandler{
		service: s,
	}

}

type UserDTO struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Login     string  `json:"login" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Balance   float64 `json:"balance"`
}

func ConvertUserToDTO(u *domain.User) UserDTO {

	return UserDTO{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Login:     u.Login,
		Password:  u.Password,
		Balance:   u.Balance,
	}

}

func (h *userHandler) AllUsersInfo(c *gin.Context) {

	slog.Info("Handler started \"AllUsersInfo\"")

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	start := time.Now()

	users, err := h.service.AllUsersInfo(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(200, nil)
		} else {
			c.JSON(500, domain.ErrorResponse{
				Error: err.Error(),
				Code:  500,
			})
		}
		return
	}

	dtos := []UserDTO{}
	for _, user := range *users {

		dtos = append(dtos, ConvertUserToDTO(&user))

	}

	if time.Since(start) > 4*time.Second {
		slog.Warn("\"AllUsersInfo\" took a lot of time")
	}

	slog.Info("Handler ended \"AllUsersInfo\" success")
	c.JSON(200, dtos)

}

func (h *userHandler) UserInfo(c *gin.Context) {

	slog.Info("Handler started \"UserInfo\"")

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	start := time.Now()

	login := c.Param("login")
	user, err := h.service.UserInfo(ctx, login)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(404, domain.ErrorResponse{
				Error: domain.ErrUserNotFound.Error(),
				Code:  404,
			})
		} else {
			c.JSON(500, domain.ErrorResponse{
				Error: err.Error(),
				Code:  500,
			})
		}
		return
	}

	dto := ConvertUserToDTO(user)

	if time.Since(start) > 4*time.Second {
		slog.Warn("\"UserInfo\" took a lot of time")
	}

	slog.Info("Handler ended \"UserInfo\" success")
	c.JSON(200, dto)

}

func (h *userHandler) RegisterAccount(c *gin.Context) {

	slog.Info("Handler started \"RegisterAccount\"")

	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	start := time.Now()

	var req UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Handler \"RegisterAccount\" get next error when parsing json:", err)
		c.JSON(400, domain.ErrorResponse{
			Error: err.Error(),
			Code:  400,
		})
		return
	}

	user, err := h.service.RegisterAccount(ctx, req.FirstName, req.LastName, req.Login, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidFirstName) || errors.Is(err, domain.ErrInvalidLastName) || errors.Is(err, domain.ErrInvalidLogin) || errors.Is(err, domain.ErrInvalidPassword) {
			c.JSON(400, domain.ErrorResponse{
				Error: err.Error(),
				Code:  400,
			})
		} else {
			c.JSON(500, domain.ErrorResponse{
				Error: err.Error(),
				Code:  500,
			})
		}
		return
	}

	dto := ConvertUserToDTO(user)

	if time.Since(start) > 4*time.Second {
		slog.Warn("\"RegisterAccount\" took a lot of time")
	}

	slog.Info("Handler ended \"RegisterAccount\" success")
	c.JSON(201, dto)

}
