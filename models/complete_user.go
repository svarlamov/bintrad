package models

type CompleteUser struct {
	Id              int64   `json:"id"`
	Username        string  `json:"username"`
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	Grade           int64   `json:"grade"`
	StudentId       int64   `json:"studentId"`
	Email           string  `json:"email"`
	StartingBalance float64 `json:"startingBalance"`
	CurrentBalance  float64 `json:"currentBalance"`
	TotalPnL        float64 `json:"totalPnL"`
	AveragePnL      float64 `json:"averagePnL"`
}

func (completeUser *CompleteUser) PopulateFromUser(user User) error {
	completeUser.Id = user.Id
	completeUser.Username = user.Username
	completeUser.FirstName = user.FirstName
	completeUser.LastName = user.LastName
	completeUser.Grade = user.Grade
	completeUser.StudentId = user.StudentId
	completeUser.Email = user.Email
	completeUser.StartingBalance = user.StartingBalance

	var contracts []Contract
	err := db.Where("user_id = ?", user.Id).Scan(&contracts).Error
	if err != nil {
		return err
	}
	var currentBalance, pnlSum float64
	currentBalance = completeUser.StartingBalance
	for _, contract := range contracts {
		currentBalance -= contract.Bet
		pnlSum += ((contract.Return / contract.Bet) - 1) * 100.0
		currentBalance += contract.Return
	}
	completeUser.CurrentBalance = currentBalance
	completeUser.TotalPnL = ((currentBalance / completeUser.StartingBalance) - 1) * 100.0
	completeUser.AveragePnL = pnlSum / float64(len(contracts))
	return nil
}
