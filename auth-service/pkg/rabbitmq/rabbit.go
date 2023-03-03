package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn      *amqp.Connection
	Publisher *amqp.Channel
}

func NewRabbit(url string) (*Rabbit, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	publisher, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	rabbit := &Rabbit{
		Conn:      conn,
		Publisher: publisher,
	}

	err = rabbit.connectToAuthTopic()
	if err != nil {
		return nil, err
	}

	return rabbit, nil
}

func (rabbit *Rabbit) Close() {
	if rabbit != nil {
		rabbit.Close()
	}

	if rabbit.Publisher != nil {
		rabbit.Publisher.Close()
	}
}
