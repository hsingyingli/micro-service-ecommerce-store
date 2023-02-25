package db

import "context"

const getProduct = `
  SELECT id, uid, title, price, amount, description, image_data, image_name, image_type, created_at, updated_at
  FROM products 
  WHERE id = $1
`

func (store *Store) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := store.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.Title,
		&i.Price,
		&i.Amount,
		&i.Description,
		&i.ImageData,
		&i.ImageName,
		&i.ImageType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProducts = `
  SELECT id, uid, title, price, amount, description, image_data, image_name, image_type, created_at, updated_at
  FROM products 
  WHERE uid = $1
  ORDER BY created_at
  LIMIT $2 
  OFFSET $3
`

func (store *Store) ListProducts(ctx context.Context, uid int64, limit int64, offset int64) ([]Product, error) {
	rows, err := store.db.QueryContext(ctx, listProducts, uid, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.UID,
			&i.Title,
			&i.Price,
			&i.Amount,
			&i.Description,
			&i.ImageData,
			&i.ImageName,
			&i.ImageType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createProduct = `
  INSERT INTO products (
    uid, title, price, amount, description, image_data, image_name, image_type 
  ) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
  )
  RETURNING id, uid, title, price, amount, description, image_data, image_name, image_type, created_at, updated_at
`

type CreateProductParam struct {
	UID         int64
	Title       string
	Price       int64
	Amount      int64
	Description string
	ImageData   []byte
	ImageName   string
	ImageType   string
}

func (store *Store) CreateProduct(ctx context.Context, args CreateProductParam) (Product, error) {
	row := store.db.QueryRowContext(ctx, createProduct,
		args.UID,
		args.Title,
		args.Price,
		args.Amount,
		args.Description,
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
		&product.Description,
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
