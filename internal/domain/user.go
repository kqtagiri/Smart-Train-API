package domain

import "strings"

type User struct {
	FirstName string
	LastName  string
	Login     string
	Password  string //8-24
	Balance   float64
}

func NewUser(firstName, lastName, login, pass string) (*User, error) {

	if firstName == "" || strings.ContainsAny(firstName, "0123456789!@#$%^&*()=+_/.,'[];") {
		return nil, ErrInvalidFirstName
	}

	if lastName == "" || strings.ContainsAny(lastName, "0123456789!@#$%^&*()=+_/.,'[];") {
		return nil, ErrInvalidLastName
	}

	if login == "" {
		return nil, ErrInvalidLogin
	}

	if pass == "" || len([]rune(pass)) < 8 || len([]rune(pass)) > 24 {
		return nil, ErrInvalidPassword
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Login:     login,
		Password:  pass,
		Balance:   0,
	}, nil

}

func (u *User) ReplenishBalance(money float64) error {

	if money <= 0 {
		return ErrInvalidReplenish
	}

	u.Balance += money
	return nil

}
