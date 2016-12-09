package models

import "time"

type TickerData struct {
	Id        int64
	TickerId  int64
	OpensAt   time.Time
	ClosesAt  time.Time
	Period    int64
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
