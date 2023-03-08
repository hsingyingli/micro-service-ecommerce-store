package rabbitmq

import (
	"cart/pkg/db"
	"encoding/json"
	"log"
)

func (rabbit *Rabbit) ListenOnAuth() error {
	err := rabbit.Consumer.ExchangeDeclare(
		"auth_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
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
		q.Name,       // queue name
		"user.*",     // routing key
		"auth_topic", // exchange
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
			var user UserPayload
			err = json.Unmarshal(d.Body, &user)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("Message receive: " + d.RoutingKey)

			switch d.RoutingKey {
			case "user.create":
				err = rabbit.ConsumeCreateUser(user)
			case "user.update":
				err = rabbit.ConsumeUpdateUser(user)
			case "user.delete":
				err = rabbit.ConsumeDeleteUser(user.ID)
			}
			log.Println(err)
		}
	}()

	return err
}

func (rabbit *Rabbit) ListenOnProduct() error {
	err := rabbit.Consumer.ExchangeDeclare(
		"product_topic", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
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
		q.Name,          // queue name
		"product.*.#",   // routing key
		"product_topic", // exchange
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
			var product []db.ProductPayload
			err = json.Unmarshal(d.Body, &product)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("Message receive: " + d.RoutingKey)

			switch d.RoutingKey {
			case "product.create":
				err = rabbit.ConsumeCreateProduct(product[0])
			case "product.delete":
				err = rabbit.ConsumeDeleteProduct(product[0])
			case "product.update.amount":
				err = rabbit.ConsumeUpdateProductAmount(product)
			case "product.update.info":
				err = rabbit.ConsumeUpdateProductInfo(product[0])
			case "product.update.infoWithImage":
				err = rabbit.ConsumeUpdateProductInfoWithImage(product[0])
			}
			log.Println(err)
		}
	}()

	return err
}
