package rabbitmq

func (rabbit *Rabbit) connectToProductTopic() error {
	err := rabbit.Publisher.ExchangeDeclare(
		"product_topic", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	return err
}
