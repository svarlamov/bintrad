package models

type CreateContractSessionResponse struct {
	Id         int64
	Price      float64
	Ttl        int64
	Period     int64
	TickerData []TickerData
	TickerType string
}
