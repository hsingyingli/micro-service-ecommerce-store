package rabbitmq

import (
	"context"
	"payment/pkg/db"
	"time"
)

func (rabbit *Rabbit) ConsumeCreateOrder(order db.OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var price int64 = 0

	for _, i := range order.Items {
		price += i.Price * i.Amount
	}

	defer cancel()
	err := rabbit.store.CreateOrder(ctx, db.CreateOrderParam{
		ID:    order.ID,
		UID:   order.UID,
		Price: price,
	})

	return err
}

func (rabbit *Rabbit) ConsumeDeleteOrder(order db.OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.DeleteOrder(ctx, order.ID, order.UID)
	return err
}
