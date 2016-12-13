package models

type CreateContractSessionResponse struct {
	Id         int64        `json:"id"`
	Price      float64      `json:"price"`
	Ttl        int64        `json:"ttl"`
	Period     int64        `json:"period"`
	TickerData []TickerData `json:"tickerData"`
	TickerType string       `json:"tickerType"`
}
