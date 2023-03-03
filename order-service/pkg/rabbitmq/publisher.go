package rabbitmq

func (rabbit *Rabbit) connectToOrderTopic() error {
	err := rabbit.Publisher.ExchangeDeclare(
		"order_topic", // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	return err
}
