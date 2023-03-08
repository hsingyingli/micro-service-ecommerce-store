package rabbitmq

import (
	"context"
	"order/pkg/db"
	"time"
)

func (rabbit *Rabbit) ConsumePaymentSuccess(order db.OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rabbit.store.UpdateOrderStatus(ctx, order.ID, "FINISH")
	if err != nil {
		return err
	}

	payload, err := rabbit.store.GetOrderInfo(ctx, order.ID, order.UID)
	err = rabbit.PublishOrder(ctx, "order.finish", payload)
	return err
}
