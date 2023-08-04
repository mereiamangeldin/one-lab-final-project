package entity

type Item struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Cost        float64 `json:"cost" db:"cost"`
}
