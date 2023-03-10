package rabbitmq

import (
	"context"
	"product/pkg/db"
	"time"
)

func (rabbit *Rabbit) ConsumeCreateOrder(order db.OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	product, err := rabbit.store.UpdateInventoryStatuTx(ctx, order, -1, 1)
	if err != nil {
		return err
	}
	err = rabbit.PublishProduct(ctx, "product.update.amount", product)
	return err
}

func (rabbit *Rabbit) ConsumeDeleteOrder(order db.OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	product, err := rabbit.store.UpdateInventoryStatuTx(ctx, order, 1, -1)
	if err != nil {
		return err
	}
	err = rabbit.PublishProduct(ctx, "product.update.amount", product)
	return err
}

func (rabbit *Rabbit) ConsumeFinishOrder(order db.OrderPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rabbit.store.UpdateInventoryStatuTx(ctx, order, 0, -1)
	return err
}
