package models

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"strings"
)

var (
	incorrectPasskeyError = errors.New("Incorrect password")
)

type AuthRequest struct {
	Username string `json:"username" valid:"required"`
	Passkey  string `json:"passkey" valid:"required"`
}

func (request *AuthRequest) Parameters() error {
	_, err := valid.ValidateStruct(request)
	if err != nil {
		return err
	}
	request.Username = strings.Split(request.Username, "@")[0]
	return nil
}

func (request *AuthRequest) Authenticate() (User, error) {
	user := User{Username: request.Username}
	err := user.FindByUsername()
	if err != nil {
		return user, err
	}
	if user.PasskeySalt != "" {
		// TODO: Handle encrypted passkeys
		return user, incorrectPasskeyError
	}
	if user.Passkey == request.Passkey {
		return user, nil
	} else {
		return user, incorrectPasskeyError
	}
}
