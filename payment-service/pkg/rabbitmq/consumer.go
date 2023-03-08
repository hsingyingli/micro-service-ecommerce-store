package rabbitmq

import (
	"encoding/json"
	"log"
	"payment/pkg/db"
)

func (rabbit *Rabbit) listenOnOrder() error {
	err := rabbit.Consumer.ExchangeDeclare(
		"order_topic", // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // interna
		false,         // no-wait
		nil,           // arguments
	)

	q, err := rabbit.Consumer.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = rabbit.Consumer.QueueBind(
		q.Name,        // queue name
		"order.*",     // routing key
		"order_topic", // exchange
		false,
		nil)

	msgs, err := rabbit.Consumer.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	go func() {
		for d := range msgs {
			var err error
			var order db.OrderPayload
			err = json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("Message receive: " + d.RoutingKey)

			switch d.RoutingKey {
			case "order.create":
				rabbit.ConsumeCreateOrder(order)
			case "order.delete":
				err = rabbit.ConsumeDeleteOrder(order)
			}
			log.Println(err)
		}
	}()

	return err
}
