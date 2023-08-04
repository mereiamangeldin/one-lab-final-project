package entity

import "time"

type Transactions struct {
	Id        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	ItemID    int64     `json:"item_id" db:"item_id"`
	Amount    float64   `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
