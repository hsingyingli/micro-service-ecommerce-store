package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn      *amqp.Connection
	Publisher *amqp.Channel
}

func NewRabbit(url string) (rabbit *Rabbit, err error) {
	rabbit = &Rabbit{}

	rabbit.Conn, err = amqp.Dial(url)

	if err != nil {
		return
	}

	rabbit.Publisher, err = rabbit.Conn.Channel()
	if err != nil {
		return
	}

	err = rabbit.connectToAuthTopic()

	return
}

func (rabbit *Rabbit) Close() {
	if rabbit != nil {
		rabbit.Close()
	}

	if rabbit.Publisher != nil {
		rabbit.Publisher.Close()
	}
}
