package db

import (
	"context"
	"time"
)

const getProductAmountById = `
  SELECT amount 
  FROM products
  WHERE id = $1 
`

func (store *Store) GetProductAmountById(ctx context.Context, id int64) (int64, error) {
	row := store.db.QueryRowContext(ctx, getProductAmountById, id)
	var amount int64
	err := row.Scan(&amount)
	return amount, err
}

const createProduct = `
  INSERT INTO products (
    id, uid, title, price, amount, image_data, image_name, image_type 
  ) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
  )
  RETURNING id, uid, title, price, amount, image_data, image_name, image_type, created_at, updated_at
`

type CreateProductParam struct {
	ID        int64
	UID       int64
	Title     string
	Price     int64
	Amount    int64
	ImageData []byte
	ImageName string
	ImageType string
}

func (store *Store) CreateProduct(ctx context.Context, args CreateProductParam) (Product, error) {
	row := store.db.QueryRowContext(ctx, createProduct,
		args.ID,
		args.UID,
		args.Title,
		args.Price,
		args.Amount,
		args.ImageData,
		args.ImageName,
		args.ImageType)

	var product Product
	err := row.Scan(
		&product.ID,
		&product.UID,
		&product.Title,
		&product.Price,
		&product.Amount,
		&product.ImageData,
		&product.ImageName,
		&product.ImageType,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return product, err
}

const updateProductAmount = `
  UPDATE products
  SET amount = $2,
      updated_at = $3
  WHERE id = $1
`

func (store *Store) UpdateProductAmount(ctx context.Context, id int64, amount int64) error {
	_, err := store.db.ExecContext(ctx, updateProductAmount, id, amount, time.Now())
	return err
}

const updateProductInfo = `
  UPDATE products
  SET title = $2,
      price = $3,
      amount = $4,
      updated_at = $5,
  WHERE id = $1
`

type UpdateProductInfoParam struct {
	ID     int64
	Title  string
	Price  int64
	Amount int64
}

func (store *Store) UpdateProductInfo(ctx context.Context, args UpdateProductInfoParam) error {
	_, err := store.db.ExecContext(ctx, updateProductInfo,
		args.ID,
		args.Title,
		args.Price,
		args.Amount,
	)
	return err
}

const updateProductInfoWithImage = `
  UPDATE products
  SET title = $2,
      price = $3,
      amount = $4,
      image_data = $5,
      image_name = $6,
      image_type = $7,
      updated_at = $8
  WHERE id = $1
`

type UpdateProductInfoWithImageParam struct {
	ID        int64
	Title     string
	Price     int64
	Amount    int64
	ImageData []byte
	ImageName string
	ImageType string
}

func (store *Store) UpdateProductInfoWithImage(ctx context.Context, args UpdateProductInfoWithImageParam) error {
	_, err := store.db.ExecContext(ctx, updateProductInfoWithImage,
		args.ID,
		args.Title,
		args.Price,
		args.Amount,
		args.ImageData,
		args.ImageName,
		args.ImageType,
		time.Now(),
	)
	return err
}

const deleteProductById = `
  DELETE 
  FROM products 
  WHERE id = $1 AND uid = $2
`

func (store *Store) DeleteProductById(ctx context.Context, id int64, uid int64) error {
	_, err := store.db.ExecContext(ctx, deleteProductById, id, uid)
	return err
}
