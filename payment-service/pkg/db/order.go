package db

import "context"

const getOrderById = `
  SELECT id, uid
  FROM orders 
  WHERE id = $1 and uid = $2
`

func (store *Store) GetOrderById(ctx context.Context, id int64, uid int64) (OrderPayload, error) {
	row := store.db.QueryRowContext(ctx, getOrderById, id, uid)
	var order OrderPayload
	err := row.Scan(&order.ID, &order.UID)
	return order, err
}

const createOrder = `
  INSERT INTO orders (
    id, uid, price
  ) VALUES (
    $1, $2, $3
  )
  `

type CreateOrderParam struct {
	ID    int64
	UID   int64
	Price int64
}

func (store *Store) CreateOrder(ctx context.Context, args CreateOrderParam) error {
	_, err := store.db.ExecContext(ctx, createOrder, args.ID, args.UID, args.Price)
	return err
}

const deleteOrder = `
  DELETE 
  FROM orders 
  WHERE id = $1 and uid = $2
`

func (store *Store) DeleteOrder(ctx context.Context, id int64, uid int64) error {
	_, err := store.db.ExecContext(ctx, deleteOrder, id, uid)
	return err
}
