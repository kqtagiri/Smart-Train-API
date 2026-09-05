package domain

import "errors"

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

var (
	//User
	ErrInvalidFirstName = errors.New("Invalid first name")
	ErrInvalidLastName  = errors.New("Invalid last name")
	ErrInvalidLogin     = errors.New("Invalid login")
	ErrInvalidPassword  = errors.New("Invalid password")

	ErrUserNotFound     = errors.New("User not found")
	ErrInvalidReplenish = errors.New("Invalid replenish")
)
