package rabbitmq

import (
	"context"
	"encoding/json"
	"product/pkg/db"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"product_topic", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		ch: ch,
	}

	return publisher, nil
}

type ProductPayload struct {
	ID        int64
	UID       int64
	Title     string
	Price     int64
	Amount    int64
	ImageData []byte
	ImageName string
	ImageType string
}

func (publisher *Publisher) ProductCreate(ctx context.Context, product db.Product) error {
	body, err := json.Marshal(ProductPayload{
		ID:        product.ID,
		UID:       product.UID,
		Title:     product.Title,
		Price:     product.Price,
		Amount:    product.Amount,
		ImageData: product.ImageData,
		ImageName: product.ImageName,
		ImageType: product.ImageType,
	})

	if err != nil {
		return nil
	}

	err = publisher.ch.PublishWithContext(ctx,
		"product_topic",  // exchange
		"product.create", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err
}

func (publisher *Publisher) ProductUpdate(ctx context.Context, product db.Product) error {
	body, err := json.Marshal(ProductPayload{
		ID:        product.ID,
		UID:       product.UID,
		Title:     product.Title,
		Price:     product.Price,
		Amount:    product.Amount,
		ImageData: product.ImageData,
		ImageName: product.ImageName,
		ImageType: product.ImageType,
	})

	if err != nil {
		return err
	}

	err = publisher.ch.PublishWithContext(ctx,
		"product_topic",  // exchange
		"product.update", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err
}

func (publisher *Publisher) ProductDelete(ctx context.Context, id int64, uid int64) error {
	body, err := json.Marshal(ProductPayload{
		ID:  id,
		UID: uid,
	})

	if err != nil {
		return err
	}

	err = publisher.ch.PublishWithContext(ctx,
		"product_topic",  // exchange
		"product.delete", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err

}
