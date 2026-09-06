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
	UserInfo(ctx context.Context, login string) (*domain.User, error)
	RegisterAccount(ctx context.Context, u *domain.User) error
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

func (r *userRepo) UserInfo(ctx context.Context, login string) (*domain.User, error) {

	slog.Info("Repository started \"UserInfo\"")

	query := `SELECT * FROM users WHERE login = $1;`
	var model UserModel
	if err := r.db.Pool.QueryRow(ctx, query, login).Scan(&model.Id, &model.FirstName, &model.LastName, &model.Login, &model.Password, &model.Balance); err != nil {
		slog.Error("Repository \"UserInfo\" get next error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	user := ConvertModelToUser(&model)

	slog.Info("Repository ended \"UserInfo\" success")
	return &user, nil

}

func (r *userRepo) RegisterAccount(ctx context.Context, u *domain.User) error {

	slog.Info("Repository started \"RegisterAccount\"")

	query := `INSERT INTO users (first_name, last_name, login, password, balance) VALUES ($1,$2,$3,$4,$5);`
	result, err := r.db.Pool.Exec(ctx, query, u.FirstName, u.LastName, u.Login, u.Password, u.Balance)
	if err != nil {
		slog.Error("Repository \"RegisterAccount\" get next error:", err)
		return err
	}

	affected := result.RowsAffected()
	if affected != 1 {
		slog.Error("Repository \"RegisterAccount\" get next error:", domain.ErrWithInsert)
		return domain.ErrWithInsert
	}

	slog.Info("Repository ended \"RegisterAccount\" success")
	return nil

}
