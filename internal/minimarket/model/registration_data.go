package model

import (
	"crypto/md5"
	"encoding/hex"
)

type RegistrationData struct {
	Login        string `json:"login"         db:"login"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}

type LoginForm struct {
	Login    string `json:"login"     validate:"min=3,max=40"`
	Password string `json:"password"  validate:"min=3,max=40"`
}

func (x LoginForm) ToRegistrationData() RegistrationData {
	hash := md5.Sum([]byte(x.Password))

	return RegistrationData{
		Login:        x.Login,
		PasswordHash: hex.EncodeToString(hash[:]),
	}
}
