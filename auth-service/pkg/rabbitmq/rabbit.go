package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn      *amqp.Connection
	Publisher *Publisher
}

func NewRabbit(url string) (*Rabbit, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		log.Fatal(err)
	}

	publisher, err := NewPublisher(conn)
	if err != nil {
		return nil, err
	}

	rabbit := &Rabbit{
		Conn:      conn,
		Publisher: publisher,
	}

	return rabbit, nil
}

func (rabbit *Rabbit) Close() {
	if rabbit != nil {
		rabbit.Close()
	}

	if rabbit.Publisher != nil && rabbit.Publisher.ch != nil {
		rabbit.Publisher.ch.Close()
	}
}
