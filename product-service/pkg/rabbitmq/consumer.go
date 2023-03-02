package rabbitmq

import (
	"product/pkg/db"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	ch    *amqp.Channel
	store *db.Store
}

func NewConsumer(conn *amqp.Connection, store *db.Store) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		ch:    ch,
		store: store,
	}

	return consumer, nil
}
