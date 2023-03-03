package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
)

type UserPayload struct {
	ID       int64
	Username string
	Email    string
}

func (rabbit *Rabbit) PublishUser(ctx context.Context, key string, user UserPayload) error {
	body, err := json.Marshal(user)

	if err != nil {
		return err
	}

	err = rabbit.Publisher.PublishWithContext(ctx,
		"auth_topic", // exchange
		key,          // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err

}
