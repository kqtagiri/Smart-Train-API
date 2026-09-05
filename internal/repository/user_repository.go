package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"smarttrain/internal/database"
	"smarttrain/internal/domain"
)

type UserRepository interface {
	AllUsersInfo(ctx context.Context) (*[]domain.User, error)
}

type userRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {

	return &userRepo{
		db: db,
	}

}

type UserModel struct {
	Id        int
	FirstName string
	LastName  string
	Login     string
	Password  string //8-24
	Balance   float64
}

func ConvertModelToUser(m *UserModel) domain.User {

	return domain.User{
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Login:     m.Login,
		Password:  m.Password,
		Balance:   m.Balance,
	}

}

func (r *userRepo) AllUsersInfo(ctx context.Context) (*[]domain.User, error) {

	slog.Info("Repository started \"AllUsersInfo\"")

	query := `SELECT * FROM users;`
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		slog.Error("Repository \"AllUsersInfo\" get next error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var model UserModel
	users := []domain.User{}
	for rows.Next() {

		if err := rows.Scan(&model.Id, &model.FirstName, &model.LastName, &model.Login, &model.Password, &model.Balance); err != nil {
			slog.Error("Repository \"AllUsersInfo\" get next error:", err)
			return nil, err
		}

		users = append(users, ConvertModelToUser(&model))

	}

	if err := rows.Err(); err != nil {
		slog.Error("Repository \"AllUsersInfo\" get next error:", err)
		return nil, err
	}

	slog.Info("Repository ended \"AllUsersInfo\" success")
	return &users, nil

}
