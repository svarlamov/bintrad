package models

import "time"

type Ticker struct {
	Id        int64
	Ticker    string
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Ticker
}
