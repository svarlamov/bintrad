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

func (ticker *Ticker) Create() error {
	return db.Create(&ticker).Error
}

func (ticker *Ticker) GetRandomTickerAndDataSubset(contractPeriod int) ([]TickerData, int64, error) {
	var ticks []TickerData
	var finalTickId int64
	// TODO: Generate random ID
	randomId := 1
	err := db.Where("id = ?", randomId).First(&ticker).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	var firstTick, lastTick TickerData
	err = db.Where("ticker_id = ?", ticker.Id).Order("opens_at ASC").First(&firstTick).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	err = db.Where("ticker_id = ?", ticker.Id).Order("opens_at DESC").First(&lastTick).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	// TODO: Generate random start time and periods count
	// randomStartTime < minBufferToTickerDataEnd
	// randomPeriods >= contractPeriod+1 && randomPeriods < periodsBuffer
	randomStartTime := firstTick.OpensAt
	randomPeriods := 3
	err = db.Where("ticker_id = ? AND opens_at > ?", ticker.Id, randomStartTime).Order("opens_at ASC").Limit(randomPeriods).Scan(&ticks).Error
	if err != nil {
		return ticks, finalTickId, err
	}
	return ticks[:len(ticks)-contractPeriod], finalTickId, nil
}
