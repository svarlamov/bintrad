package models

import "time"

type ContractSession struct {
	Id        int64
	UserId    int64
	TickerId  int64
	Price     float64
	Ttl       int64
	Period    int64
	DataStart time.Time
	DataEnd   time.Time
	IsClosed  bool
	ClosedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
