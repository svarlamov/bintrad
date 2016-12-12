package models

import "time"

type TickerData struct {
	Id        int64     `json:"-"`
	TickerId  int64     `json:"-"`
	OpensAt   time.Time `json:"opensAt"`
	ClosesAt  time.Time `json:"closesAt"`
	Period    int64     `json:"period"`
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Close     float64   `json:"close"`
	Volume    float64   `json:"volume"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (tickerData *TickerData) Create() error {
	return db.Create(&tickerData).Error
}

func (tickerData *TickerData) FindById() error {
	return db.Where("id = ?", tickerData.Id).First(&tickerData).Error
}
