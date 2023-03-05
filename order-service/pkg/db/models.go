package db

import "time"

type Product struct {
	ID        int64     `json:"id"`
	UID       int64     `json:"uid"`
	Title     string    `json:"title"`
	Price     int64     `json:"price"`
	Amount    int64     `json:"amount"`
	ImageName string    `json:"image_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductInfo struct {
	ID     int64 `json:"id"`
	UID    int64 `json:"uid"`
	Price  int64 `json:"price"`
	Amount int64 `json:"amount"`
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID        int64     `json:"id"`
	UID       int64     `json:"uid"`
	PID       int64     `json:"pid"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderDetail struct {
	ID        int64  `json:"id"`
	PID       int64  `json:"pid"`
	Amount    int64  `json:"amount"`
	Title     string `json:"title"`
	Price     int64  `json:"price"`
	ImageName string `json:"image_name"`
}
