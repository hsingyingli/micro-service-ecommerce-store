package rabbitmq

import (
	"authentication/pkg/db"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
)

type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"auth_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		ch: ch,
	}

	return publisher, nil
}

type UserPayload struct {
	ID       int64
	Username string
	Email    string
}

func (publisher *Publisher) UserCreated(ctx context.Context, user db.User) error {

	body, err := json.Marshal(UserPayload{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		return err
	}

	err = publisher.ch.PublishWithContext(ctx,
		"auth_topic",  // exchange
		"user.create", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err
}
