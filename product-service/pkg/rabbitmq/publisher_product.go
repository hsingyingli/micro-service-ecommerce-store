package rabbitmq

import (
	"context"
	"encoding/json"
	"product/pkg/db"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) PublishProduct(ctx context.Context, key string, products []db.ProductPayload) error {
	body, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = rabbit.Publisher.PublishWithContext(ctx,
		"product_topic", // exchange
		key,             // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err

}
