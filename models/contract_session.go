package models

import "time"

type ContractSession struct {
	Id          int64
	UserId      int64
	TickerId    int64
	Price       float64
	Ttl         int64
	Period      int64
	DataStart   time.Time
	DataEnd     time.Time
	FinalTickId int64
	IsClosed    bool
	ClosedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (session *ContractSession) Create() error {
	return db.Create(&session).Error
}

func (session *ContractSession) GetByIdAndUserId() error {
	return db.Where("id = ? AND user_id = ?", session.Id, session.UserId).First(&session).Error
}

func (session *ContractSession) CloseByIdAndUserId() error {
	return db.Model(&ContractSession{}).Where("id = ? AND user_id = ?", session.Id, session.UserId).Updates(map[string]interface{}{"is_closed": true, "closed_at": time.Now()}).Error
}

func (session *ContractSession) GenerateContractData(bet float64, isBullish bool, finalTickId int64) (Contract, error) {
	finalTick := TickerData{
		Id: finalTickId,
	}
	err := finalTick.FindById()
	if err != nil {
		return Contract{}, err
	}
	contract := Contract{
		UserId:            session.UserId,
		TickerId:          session.TickerId,
		ContractSessionId: session.Id,
		Bet:               bet,
		Price:             session.Price,
		IsBullish:         isBullish,
	}
	if contract.IsBullish && finalTick.Open > contract.Price {
		contract.IsCorrect = true
		contract.Return = 2 * contract.Bet
	} else if contract.IsBullish && finalTick.Open < contract.Price {
		contract.IsCorrect = false
		contract.Return = 0
	} else if !contract.IsBullish && finalTick.Open < contract.Price {
		contract.IsCorrect = true
		contract.Return = 2 * contract.Bet
	} else if !contract.IsBullish && finalTick.Open > contract.Price {
		contract.IsCorrect = false
		contract.Return = 0
	}
	return contract, nil
}
