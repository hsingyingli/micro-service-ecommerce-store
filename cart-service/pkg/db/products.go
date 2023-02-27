package db

import "context"

const createProduct = `
  INSERT INTO products (
    uid, title, price, amount, description, image_data, image_name, image_type 
  ) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
  )
  RETURNING id, uid, title, price, amount, description, image_data, image_name, image_type, created_at, updated_at
`

type CreateProductParam struct {
	UID       int64
	Title     string
	Price     int64
	ImageData []byte
	ImageName string
	ImageType string
}

func (store *Store) CreateProduct(ctx context.Context, args CreateProductParam) (Product, error) {
	row := store.db.QueryRowContext(ctx, createProduct,
		args.UID,
		args.Title,
		args.Price,
		args.ImageData,
		args.ImageName,
		args.ImageType)

	var product Product
	err := row.Scan(
		&product.ID,
		&product.UID,
		&product.Title,
		&product.Price,
		&product.ImageData,
		&product.ImageName,
		&product.ImageType,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return product, err
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
