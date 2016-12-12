package models

type FinaliseContractSessionRequest struct {
	Bet       float64 `json:"bet"`
	IsBullish bool    `json:"isBullish"`
}
