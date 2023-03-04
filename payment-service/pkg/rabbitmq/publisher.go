package rabbitmq

func (rabbit *Rabbit) connectToPaymentTopic() error {
	err := rabbit.Publisher.ExchangeDeclare(
		"payment_topic", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	return err
}
