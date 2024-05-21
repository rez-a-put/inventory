package controller

import (
	m "inventory/model"
	r "inventory/repository"
	u "inventory/utils"

	"golang.org/x/crypto/bcrypt"
)

// Login : to login using email and password
func Login(email, password string) (token string, err error) {
	var (
		users []*m.User
	)

	users, err = r.GetUsers("", email, "", []string{"1"})
	if err != nil {
		return "", err
	}

	err = u.VerifyPassword(password, users[0].Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err = u.GenerateToken(users[0].ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
