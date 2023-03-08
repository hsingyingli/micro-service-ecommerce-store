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

type ProductPayload struct {
	ID        int64
	UID       int64
	Title     string
	Price     int64
	Amount    int64
	ImageName string
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
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID        int64     `json:"id"`
	OID       int64     `json:"oid"`
	PID       int64     `json:"pid"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItemDetail struct {
	ID        int64  `json:"id"`
	OID       int64  `json:"oid"`
	PID       int64  `json:"pid"`
	Amount    int64  `json:"amount"`
	Price     int64  `json:"price"`
	ImageName string `json:"image_name"`
}

type OrderPayload struct {
	ID    int64             `json:"id"`
	UID   int64             `json:"uid"`
	Items []OrderItemDetail `json:"items"`
}

type OrderInfo struct {
	PID    int64 `json:"pid"`
	Amount int64 `json:"amount"`
}
