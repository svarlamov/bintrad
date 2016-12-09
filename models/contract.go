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
