package models

type LeaderboardUser struct {
	Id              int64   `json:"-"`
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	StartingBalance float64 `json:"-"`
	BetSum          float64 `json:"-"`
	CurrentBalance  float64 `json:"currentBalance"`
}

func GetCurrentLeaderboard(limit int64) ([]LeaderboardUser, error) {
	var users []LeaderboardUser
	err := db.Raw("select users.*, (users.starting_balance + IFNULL(bet_sum,0)) as current_balance from (select u.id, u.first_name, u.last_name, u.starting_balance as starting_balance, (select sum(`return`)-sum(bet) from contract where user_id = u.id) as bet_sum from user u) as users order by current_balance desc limit ?;", limit).Scan(&users).Error
	return users, err
}
