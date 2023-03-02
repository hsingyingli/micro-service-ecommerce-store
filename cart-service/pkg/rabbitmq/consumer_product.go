package rabbitmq

import (
	"encoding/json"
	"log"
)

func (consumer *Consumer) ListenOnProduct() error {
	err := consumer.ch.ExchangeDeclare(
		"product_topic", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	q, err := consumer.ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = consumer.ch.QueueBind(
		q.Name,          // queue name
		"product.*.#",   // routing key
		"product_topic", // exchange
		false,
		nil)

	msgs, err := consumer.ch.Consume(
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
			var product ProductPayload
			err = json.Unmarshal(d.Body, &product)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("Message receive: " + d.RoutingKey)

			switch d.RoutingKey {
			case "product.create":
				consumer.CreateProduct(product)
			case "product.delete":
				err = consumer.DeleteProduct(product.ID, product.UID)
			case "product.update.amount":
				err = consumer.UpdateProductAmount(product)
			case "product.update.info":
				err = consumer.UpdateProductInfo(product)
			case "product.update.infoWithImage":
				err = consumer.UpdateProductInfoWithImage(product)
			}
		}
	}()

	return err
}
