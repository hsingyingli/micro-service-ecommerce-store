package db

import "context"

const getOrderById = `
  SELECT id, pid, uid, amount, price
  FROM orders 
  WHERE id = $1 and uid = $2
`

func (store *Store) GetOrderById(ctx context.Context, id int64, uid int64) (OrderPayload, error) {
	row := store.db.QueryRowContext(ctx, getOrderById, id, uid)
	var order OrderPayload
	err := row.Scan(&order.ID, &order.PID, &order.UID, &order.Amount, &order.Price)
	return order, err
}

const createOrder = `
  INSERT INTO orders (
    id, pid, uid, amount, price
  ) VALUES (
    $1, $2, $3, $4, $5
  )
  `

type CreateOrderParam struct {
	ID     int64
	PID    int64
	UID    int64
	Amount int64
	Price  int64
}

func (store *Store) CreateOrder(ctx context.Context, args CreateOrderParam) error {
	_, err := store.db.ExecContext(ctx, createOrder, args.ID, args.PID, args.UID, args.Amount, args.Price)
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
