package db

import "context"

const getProduct = `
  SELECT id, uid, title, price, amount, num_unpaid, description, image_name, created_at, updated_at
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
		&i.NumUnPaid,
		&i.Description,
		&i.ImageName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProducts = `
  SELECT id, uid, title, price, amount, num_unpaid, description, image_name, created_at, updated_at
  FROM products 
  ORDER BY created_at
  LIMIT $1
  OFFSET $2
`

func (store *Store) ListProducts(ctx context.Context, limit int64, offset int64) ([]Product, error) {
	rows, err := store.db.QueryContext(ctx, listProducts, limit, offset)
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
			&i.NumUnPaid,
			&i.Description,
			&i.ImageName,
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

const listOwnProducts = `
  SELECT id, uid, title, price, amount, num_unpaid, description, image_name, created_at, updated_at
  FROM products 
  WHERE uid = $1
  ORDER BY created_at
  LIMIT $2 
  OFFSET $3
`

func (store *Store) ListOwnProducts(ctx context.Context, uid int64, limit int64, offset int64) ([]Product, error) {
	rows, err := store.db.QueryContext(ctx, listOwnProducts, uid, limit, offset)
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
			&i.NumUnPaid,
			&i.Description,
			&i.ImageName,
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
    uid, title, price, amount, description, image_name 
  ) VALUES (
    $1, $2, $3, $4, $5, $6
  )
  RETURNING id, uid, title, price, amount, num_unpaid, description, image_name, created_at, updated_at
`

type CreateProductParam struct {
	UID         int64
	Title       string
	Price       int64
	Amount      int64
	Description string
	ImageName   string
}

func (store *Store) CreateProduct(ctx context.Context, args CreateProductParam) (Product, error) {
	row := store.db.QueryRowContext(ctx, createProduct,
		args.UID,
		args.Title,
		args.Price,
		args.Amount,
		args.Description,
		args.ImageName,
	)

	var product Product
	err := row.Scan(
		&product.ID,
		&product.UID,
		&product.Title,
		&product.Price,
		&product.Amount,
		&product.NumUnPaid,
		&product.Description,
		&product.ImageName,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return product, err
}

const updateProductInfo = `

`

const deleteProductById = `
  DELETE 
  FROM products 
  WHERE id = $1 AND uid = $2
`

func (store *Store) DeleteProductById(ctx context.Context, id int64, uid int64) error {
	_, err := store.db.ExecContext(ctx, deleteProductById, id, uid)
	return err
}
