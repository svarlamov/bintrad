package models

import "time"

type User struct {
	Id          int64
	Username    string
	FirstName   string
	LastName    string
	MiddleName  string
	Grade       int64
	Gender      string
	Passkey     string
	PasskeySalt string
	StudentId   int64
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (user *User) FindById() error {
	return db.Where("id = ?", user.Id).First(&user).Error
}

func (user *User) FindByUsername() error {
	return db.Where("username = ?", user.Username).First(&user).Error
}

func (user *User) GenerateAndPersistAccessToken() (AccessToken, error) {
	at := AccessToken{
		UserId: user.Id,
		Token:  GetAccessTokenString(),
	}
	err := at.Create()
	return at, err
}

func (user *User) Create() error {
	return db.Create(&user).Error
}
