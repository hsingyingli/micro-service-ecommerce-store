package db

import (
	"context"
	"time"
)

const getProductInfo = `
  SELECT id, uid, price, amount 
  FROM products 
  WHERE id = $1
`

func (store *Store) GetProductInfo(ctx context.Context, id int64) (ProductInfo, error) {
	row := store.db.QueryRowContext(ctx, getProductInfo, id)
	var product ProductInfo
	err := row.Scan(&product.ID, &product.UID, &product.Price, &product.Amount)
	return product, err
}

const getProductPriceById = `
  SELECT price 
  FROM products 
  WHERE id = $1
`

func (store *Store) GetProductPriceById(ctx context.Context, id int64) (int64, error) {
	row := store.db.QueryRowContext(ctx, getProductPriceById, id)
	var price int64
	err := row.Scan(&price)
	return price, err
}

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
    id, uid, title, price, amount,  image_name 
  ) VALUES (
    $1, $2, $3, $4, $5, $6
  )
  RETURNING id, uid, title, price, amount, image_name, created_at, updated_at
`

type CreateProductParam struct {
	ID        int64
	UID       int64
	Title     string
	Price     int64
	Amount    int64
	ImageName string
}

func (store *Store) CreateProduct(ctx context.Context, args CreateProductParam) (Product, error) {
	row := store.db.QueryRowContext(ctx, createProduct,
		args.ID,
		args.UID,
		args.Title,
		args.Price,
		args.Amount,
		args.ImageName)

	var product Product
	err := row.Scan(
		&product.ID,
		&product.UID,
		&product.Title,
		&product.Price,
		&product.Amount,
		&product.ImageName,
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
      image_name = $5,
      updated_at = $6
  WHERE id = $1
`

type UpdateProductInfoWithImageParam struct {
	ID        int64
	Title     string
	Price     int64
	Amount    int64
	ImageName string
}

func (store *Store) UpdateProductInfoWithImage(ctx context.Context, args UpdateProductInfoWithImageParam) error {
	_, err := store.db.ExecContext(ctx, updateProductInfoWithImage,
		args.ID,
		args.Title,
		args.Price,
		args.Amount,
		args.ImageName,
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
