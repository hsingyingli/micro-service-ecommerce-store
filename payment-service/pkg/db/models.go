package db

import "time"

type Order struct {
	ID        int64     `json:"id"`
	PID       int64     `json:"pid"`
	UID       int64     `json:"uid"`
	Amount    int64     `json:"amount"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payment struct {
	ID        int64     `json:"id"`
	OID       int64     `json:"oid"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItemDetail struct {
	ID        int64  `json:"id"`
	OID       int64  `json:"oid"`
	PID       int64  `json:"pid"`
	Amount    int64  `json:"amount"`
	Price     int64  `json:"price"`
	Title     string `json:"title"`
	ImageName string `json:"image_name"`
}

type OrderPayload struct {
	ID    int64             `json:"id"`
	UID   int64             `json:"uid"`
	Items []OrderItemDetail `json:"items"`
}
