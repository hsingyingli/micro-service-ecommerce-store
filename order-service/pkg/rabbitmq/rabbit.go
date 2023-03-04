package rabbitmq

import (
	"order/pkg/db"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Conn      *amqp.Connection
	Publisher *amqp.Channel
	Consumer  *amqp.Channel
	store     *db.Store
}

func NewRabbit(url string, store *db.Store) (rabbit *Rabbit, err error) {
	rabbit = &Rabbit{}
	rabbit.store = store
	rabbit.Conn, err = amqp.Dial(url)
	if err != nil {
		return
	}

	rabbit.Publisher, err = rabbit.Conn.Channel()
	if err != nil {
		return
	}

	rabbit.Consumer, err = rabbit.Conn.Channel()
	if err != nil {
		return
	}

	err = rabbit.ListenOnAuth()
	if err != nil {
		return
	}
	err = rabbit.ListenOnProduct()
	if err != nil {
		return
	}
	err = rabbit.ListenOnPayment()
	if err != nil {
		return
	}
	err = rabbit.connectToOrderTopic()

	return
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
