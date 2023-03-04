package rabbitmq

import (
	"context"
	"time"
)

type OrderPayload struct {
	ID     int64
	PID    int64
	UID    int64 // product buyer
	CID    int64
	Amount int64
	Price  int64
}

func (rabbit *Rabbit) ConsumeCreateOrder(order OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.DeleteCart(ctx, order.CID, order.UID)
	return err
}
