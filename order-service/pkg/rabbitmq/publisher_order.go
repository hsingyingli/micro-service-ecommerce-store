package rabbitmq

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderPayload struct {
	ID     int64
	PID    int64
	UID    int64 // product buyer
	CID    int64
	Amount int64
	Price  int64
}

func (rabbit *Rabbit) PublishOrder(ctx context.Context, key string, order OrderPayload) error {
	body, err := json.Marshal(order)

	if err != nil {
		return err
	}

	err = rabbit.Publisher.PublishWithContext(ctx,
		"order_topic", // exchange
		key,           // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return err
}
