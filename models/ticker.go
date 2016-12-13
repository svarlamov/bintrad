package models

import (
	"github.com/svarlamov/bintrad/utils"
	"time"
)

type Ticker struct {
	Id        int64
	Ticker    string
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Ticker
}

func (ticker *Ticker) Create() error {
	return db.Create(&ticker).Error
}

func (ticker *Ticker) GetRandomTickerAndDataSubset(contractTickCnt int) ([]TickerData, int64, error) {
	var ticks []TickerData
	var finalTickId int64
	err := db.Model(Ticker{}).Order("RAND()").Limit(1).Scan(ticker).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	var tickCnt int64
	var firstTick, lastTick TickerData
	err = db.Where("ticker_id = ?", ticker.Id).Order("opens_at ASC").First(&firstTick).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	err = db.Where("ticker_id = ?", ticker.Id).Order("opens_at DESC").First(&lastTick).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	err = db.Model(TickerData{}).Where("ticker_id = ?", ticker.Id).Count(&tickCnt).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	var periodsBuffer int64 = 31
	offset := utils.GenerateRandomIntegerWithinRange(periodsBuffer, tickCnt-periodsBuffer)
	err = db.Where("ticker_id = ?", ticker.Id).Order("opens_at ASC").Offset(int(offset)).Limit(int(periodsBuffer)).Find(&ticks).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	finalTickId = ticks[len(ticks)-1].Id
	return ticks[:len(ticks)-contractTickCnt], finalTickId, nil
}
