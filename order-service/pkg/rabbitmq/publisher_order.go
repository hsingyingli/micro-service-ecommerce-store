package rabbitmq

import (
	"context"
	"encoding/json"
	"order/pkg/db"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) PublishOrder(ctx context.Context, key string, order db.OrderPayload) error {
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
