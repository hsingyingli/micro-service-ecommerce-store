package rabbitmq

import (
	"cart/pkg/db"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn      *amqp.Connection
	Publisher *Publisher
	Consumer  *Consumer
}

func NewRabbit(url string, store *db.Store) (*Rabbit, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := NewPublisher(conn)
	if err != nil {
		return nil, err
	}
	consumer, err := NewConsumer(conn, store)

	if err != nil {
		return nil, err
	}

	err = consumer.ListenOnAuth()
	if err != nil {
		return nil, err
	}

	err = consumer.ListenOnProduct()
	if err != nil {
		return nil, err
	}

	rabbit := &Rabbit{
		Conn:      conn,
		Publisher: publisher,
		Consumer:  consumer,
	}

	return rabbit, nil
}

func (rabbit *Rabbit) Close() {
	if rabbit.Conn != nil {
		rabbit.Conn.Close()
	}

	if rabbit.Publisher != nil && rabbit.Publisher.ch != nil {
		rabbit.Publisher.ch.Close()
	}

	if rabbit.Consumer != nil && rabbit.Consumer.ch != nil {
		rabbit.Consumer.ch.Close()
	}
}
