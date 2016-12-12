package models

import "time"

type Contract struct {
	Id                int64
	UserId            int64
	TickerId          int64
	ContractSessionId int64
	Bet               float64
	Price             float64
	IsBullish         bool
	IsCorrect         bool
	Return            float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (contract *Contract) Create() error {
	return db.Create(&contract).Error
}

func (contract *Contract) GetById() error {
	return db.Where("id = ?", contract.Id).First(&contract).Error
}
