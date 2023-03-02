package rabbitmq

import (
	"log"
	"product/pkg/db"

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

	//consumer.ListenOnAuth()

	rabbit := &Rabbit{
		Conn:      conn,
		Publisher: publisher,
		Consumer:  consumer,
	}

	return rabbit, nil
}

func (rabbit *Rabbit) Close() {
	rabbit.Conn.Close()
	rabbit.Publisher.ch.Close()
	rabbit.Consumer.ch.Close()
}
