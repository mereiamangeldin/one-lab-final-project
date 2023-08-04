package entity

type Balance struct {
	UserID int64   `json:"user_id" db:"user_id"`
	Amount float64 `json:"amount" db:"amount"`
}
