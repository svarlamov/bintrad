package models

import (
	"github.com/svarlamov/bintrad/utils"
	"time"
)

type AccessToken struct {
	Id        int64
	UserId    int64
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (token *AccessToken) Create() error {
	return db.Create(&token).Error
}

func GetAccessTokenString() string {
	return utils.GenerateRandomStringToken(128)
}

func (token *AccessToken) IsValid() bool {
	return token.Id != 0
}

func (token *AccessToken) FindByToken() error {
	return db.Where("token = ?", token.Token).First(&token).Error
}
