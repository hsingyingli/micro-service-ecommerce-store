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

type OrderPayload struct {
	ID     int64 `json:"id"`
	PID    int64 `json:"pid"`
	UID    int64 `json:"uid"`
	CID    int64 `json:"cid"`
	Amount int64 `json:"amount"`
	Price  int64 `json:"price"`
}

type Payment struct {
	ID        int64     `json:"id"`
	OID       int64     `json:"oid"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
