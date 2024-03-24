package generator

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(login string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(login), 8)
	if err != nil {
		return "", err
	}
	return string(password[:]), nil
}
