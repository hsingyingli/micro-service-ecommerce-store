package rabbitmq

import (
	"cart/pkg/db"
	"context"
	"time"
)

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

func (consumer *Consumer) CreateProduct(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := consumer.store.CreateProduct(ctx, db.CreateProductParam{
		ID:        product.ID,
		UID:       product.UID,
		Title:     product.Title,
		Price:     product.Price,
		Amount:    product.Amount,
		ImageData: product.ImageData,
		ImageName: product.ImageName,
		ImageType: product.ImageType,
	})
	return err
}

func (consumer *Consumer) UpdateProductAmount(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := consumer.store.UpdateProductAmount(ctx, product.ID, product.Amount)
	return err
}

func (consumer *Consumer) UpdateProductInfo(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := consumer.store.UpdateProductInfo(ctx, db.UpdateProductInfoParam{
		ID:     product.ID,
		Title:  product.Title,
		Price:  product.Price,
		Amount: product.Amount,
	})
	return err
}

func (consumer *Consumer) UpdateProductInfoWithImage(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := consumer.store.UpdateProductInfoWithImage(ctx, db.UpdateProductInfoWithImageParam{
		ID:        product.ID,
		Title:     product.Title,
		Price:     product.Price,
		Amount:    product.Amount,
		ImageData: product.ImageData,
		ImageName: product.ImageName,
		ImageType: product.ImageType,
	})
	return err
}

func (consumer *Consumer) DeleteProduct(id int64, uid int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := consumer.store.DeleteProductById(ctx, id, uid)
	return err
}
