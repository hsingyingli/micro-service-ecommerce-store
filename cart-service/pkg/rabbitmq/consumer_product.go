package rabbitmq

import (
	"cart/pkg/db"
	"context"
	"time"
)

func (rabbit *Rabbit) ConsumeCreateProduct(product db.ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rabbit.store.CreateProduct(ctx, db.CreateProductParam{
		ID:        product.ID,
		UID:       product.UID,
		Title:     product.Title,
		Price:     product.Price,
		Amount:    product.Amount,
		ImageName: product.ImageName,
	})
	return err
}

func (rabbit *Rabbit) ConsumeUpdateProductAmount(product []db.ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.UpdateBatchProductAmountTx(ctx, product)
	return err
}

func (rabbit *Rabbit) ConsumeUpdateProductInfo(product db.ProductPayload) error {
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

func (rabbit *Rabbit) ConsumeUpdateProductInfoWithImage(product db.ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbit.store.UpdateProductInfoWithImage(ctx, db.UpdateProductInfoWithImageParam{
		ID:        product.ID,
		Title:     product.Title,
		Price:     product.Price,
		Amount:    product.Amount,
		ImageName: product.ImageName,
	})
	return err
}

func (rabbit *Rabbit) ConsumeDeleteProduct(product db.ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rabbit.store.DeleteProductById(ctx, product.ID, product.UID)
	return err
}
