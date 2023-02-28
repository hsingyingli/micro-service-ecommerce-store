package rabbitmq

import "encoding/json"

func (consumer *Consumer) ListenOnAuth() error {
	err := consumer.ch.ExchangeDeclare(
		"auth_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
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
		q.Name,       // queue name
		"user.*",     // routing key
		"auth_topic", // exchange
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
			var user UserPayload
			err = json.Unmarshal(d.Body, &user)
			if err != nil {
				continue
			}

			switch d.RoutingKey {
			case "user.create":
				consumer.CreateUser(user)
			case "user.update":
				err = consumer.UpdateUser(user)
			case "user.delete":
				err = consumer.DeleteUser(user.ID)
			}
		}
	}()

	return err
}
