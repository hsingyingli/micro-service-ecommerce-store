package rabbitmq

import (
	"context"
	"time"
)

func (rabbit *Rabbit) ConsumePaymentSuccess(order OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.UpdateOrderStatus(ctx, order.ID, "FINISH")
	if err != nil {
		return err
	}
	err = rabbit.PublishOrder(ctx, "order.finish", order)
	return err
}
