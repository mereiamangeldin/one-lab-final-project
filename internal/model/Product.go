package model

import "time"

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	CategoryID  uint    `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Img         string  `json:"img"`
	Liked       bool    `json:"liked"`
}

type Favorite struct {
	ID        uint `json:"id"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Liked     bool `json:"liked"`
}

type Transaction struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	UserID    uint      `json:"user_id"`
	Amount    float64   `json:"amount"`
	TakenAt   time.Time `json:"taken_at"`
}

type TransactionRequest struct {
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Price     float64 `json:"price"`
}
