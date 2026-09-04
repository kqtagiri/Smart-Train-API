package repository

import "smarttrain/internal/database"

type UserRepository interface {
}

type userRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {

	return &userRepo{
		db: db,
	}

}
