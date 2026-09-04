package domain

import "errors"

var (
	//User
	ErrInvalidFirstName = errors.New("Invalid first name")
	ErrInvalidLastName  = errors.New("Invalid last name")
	ErrInvalidLogin     = errors.New("Invalid login")
	ErrInvalidPassword  = errors.New("Invalid password")

	ErrInvalidReplenish = errors.New("Invalid replenish")
)
