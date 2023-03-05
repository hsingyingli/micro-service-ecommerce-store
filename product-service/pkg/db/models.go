package db

import "time"

type Product struct {
	ID          int64     `json:"id"`
	UID         int64     `json:"uid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Amount      int64     `json:"amount"`
	NumUnPaid   int64     `json:"num_unpaid"`
	ImageName   string    `json:"image_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
