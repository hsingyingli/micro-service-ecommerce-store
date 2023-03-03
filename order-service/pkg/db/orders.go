package db

import "context"

const getOrderAmount = `
  SELECT pid, amount
  FROM orders 
  WHERE id = $1
`

func (store *Store) GetOrderAmount(ctx context.Context, id int64) (int64, int64, error) {
	row := store.db.QueryRowContext(ctx, getOrderAmount, id)
	var amount int64
	var pid int64
	err := row.Scan(&pid, &amount)
	return pid, amount, err
}

const createOrder = `
  INSERT INTO orders (
    uid, pid, amount
  ) VALUES (
    $1, $2, $3
  ) 
  RETURNING id, uid, pid, amount, created_at, updated_at
`

type CreateOrderParam struct {
	UID    int64
	PID    int64
	Amount int64
}

func (store *Store) CreateOrder(ctx context.Context, args CreateOrderParam) (Order, error) {
	row := store.db.QueryRowContext(ctx, createOrder, args.UID, args.PID, args.Amount)

	var order Order

	err := row.Scan(
		&order.ID,
		&order.UID,
		&order.PID,
		&order.Amount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	return order, err
}

const listOrders = `
  SELECT o.id, o.pid, o.amount, p.title, p.price, p.image_data, p.image_type
  FROM products as p, order as o 
  WHERE o.uid = $1
  ORDER BY o.created_at
  LIMIT $2
  OFFSET $3
`

func (store *Store) ListOrders(ctx context.Context, uid int64, limit int64, offset int64) ([]OrderDetail, error) {
	rows, err := store.db.QueryContext(ctx, listOrders, uid, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []OrderDetail

	for rows.Next() {
		var i OrderDetail
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
		orders = append(orders, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, err
}

const deleteOrder = `
  DELETE 
  FROM orders 
  WHERE id = $1 AND uid = $2
`

func (store *Store) DeleteOrder(ctx context.Context, id int64, uid int64) error {
	_, err := store.db.ExecContext(ctx, deleteOrder, id, uid)
	return err
}
