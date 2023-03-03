package rabbitmq

import (
	"context"
	"order/pkg/db"
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

func (rabbit *Rabbit) ConsumeCreateProduct(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rabbit.store.CreateProduct(ctx, db.CreateProductParam{
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

func (rabbit *Rabbit) ConsumeUpdateProductAmount(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.UpdateProductAmount(ctx, product.ID, product.Amount)
	return err
}

func (rabbit *Rabbit) ConsumeUpdateProductInfo(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.UpdateProductInfo(ctx, db.UpdateProductInfoParam{
		ID:     product.ID,
		Title:  product.Title,
		Price:  product.Price,
		Amount: product.Amount,
	})
	return err
}

func (rabbit *Rabbit) ConsumeUpdateProductInfoWithImage(product ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.UpdateProductInfoWithImage(ctx, db.UpdateProductInfoWithImageParam{
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

func (rabbit *Rabbit) ConsumeDeleteProduct(id int64, uid int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rabbit.store.DeleteProductById(ctx, id, uid)
	return err
}
