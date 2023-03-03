package rabbitmq

import (
	"log"
	"product/pkg/db"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn      *amqp.Connection
	Publisher *amqp.Channel
	Consumer  *amqp.Channel
	store     *db.Store
}

func NewRabbit(url string, store *db.Store) (*Rabbit, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	consumer, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	rabbit := &Rabbit{
		Conn:      conn,
		Publisher: publisher,
		Consumer:  consumer,
		store:     store,
	}

	err = rabbit.connectToProductTopic()
	if err != nil {
		return nil, err
	}

	err = rabbit.listenOnOrder()
	if err != nil {
		return nil, err
	}

	return rabbit, nil
}

func (rabbit *Rabbit) Close() {
	if rabbit.Conn != nil {
		rabbit.Conn.Close()
	}

	if rabbit.Publisher != nil {
		rabbit.Publisher.Close()
	}

	if rabbit.Consumer != nil {
		rabbit.Consumer.Close()
	}
}
