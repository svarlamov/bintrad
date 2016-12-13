package models

type FinaliseContractSessionResponse struct {
	Bet       float64      `json:"bet"`
	Price     float64      `json:"price"`
	IsBullish bool         `json:"isBullish"`
	IsCorrect bool         `json:"isCorrect"`
	Return    float64      `json:"return"`
	UserData  CompleteUser `json:"userData"`
}
