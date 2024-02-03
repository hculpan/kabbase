package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string    `json:"username"`
	Passkey   string    `json:"passkey"`
	LastLogin time.Time `json:"last_login"`
}

func NewUser(username, passkey string) *User {
	result := User{Username: username, LastLogin: time.Now()}
	result.SetPasskey(passkey)
	return &result
}

func (u *User) SetPasskey(passkey string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passkey), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Passkey = string(hashedPassword)
	return nil
}

func (u *User) ComparePasskey(passkey string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(u.Passkey), []byte(passkey))
}
