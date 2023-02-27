package db

import "context"

const createCart = `
  INSERT INTO carts (
    uid, pid, amount
  ) VALUES (
    $1, $2, $3
  ) 
  RETURNING id, uid, pid, amount, created_at, updated_at
`

type CreateCartParam struct {
	UID    int64
	PID    int64
	Amount int64
}

func (store *Store) CreateCart(ctx context.Context, args CreateCartParam) (Cart, error) {
	row := store.db.QueryRowContext(ctx, createCart, args.UID, args.PID, args.Amount)

	var cart Cart

	err := row.Scan(
		&cart.ID,
		&cart.UID,
		&cart.PID,
		&cart.Amount,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	return cart, err
}

const listCarts = `
  SELECT c.id, c.pid, c.amount, p.title, p.price, p.image_data, p.image_type
  FROM products as p, carts as c
  WHERE c.uid = $1
  ORDER BY c.created_at
  LIMIT $2
  OFFSET $3
`

func (store *Store) ListCarts(ctx context.Context, uid int64, limit int64, offset int64) ([]CartDetail, error) {
	rows, err := store.db.QueryContext(ctx, listCarts, uid, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []CartDetail

	for rows.Next() {
		var i CartDetail
		if err := rows.Scan(
			&i.ID,
			&i.PID,
			&i.Amount,
			&i.Title,
			&i.Price,
			&i.ImageData,
			&i.ImageType,
		); err != nil {
			return nil, err
		}
		carts = append(carts, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return carts, err
}

const updateCart = `
  UPDATE carts
  SET amount = $3
  WHERE id = $1 and uid = $2
  RETURNING id, uid, pid, amount, created_at, updated_at
`

func (store *Store) UpdateCart(ctx context.Context, id int64, uid int64, amount int64) (Cart, error) {
	row := store.db.QueryRowContext(ctx, updateCart, id, uid, amount)
	var cart Cart
	err := row.Scan(
		&cart.ID,
		&cart.UID,
		&cart.PID,
		&cart.Amount,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	return cart, err
}

const deleteCart = `
  DELETE 
  FROM carts 
  WHERE id = $1 AND uid = $2
`

func (store *Store) DeleteCart(ctx context.Context, id int64, uid int64) error {
	_, err := store.db.ExecContext(ctx, deleteCart, id, uid)
	return err
}
