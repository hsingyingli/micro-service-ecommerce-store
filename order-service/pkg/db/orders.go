package db

import "context"

const getOrderInfo = `
  SELECT oi.pid, oi.amount
  FROM order_items as oi
  WHERE oid = (
    SELECT id 
    FROM orders
    WHERE id = $1 AND uid = $2
  )
`

func (store *Store) GetOrderInfo(ctx context.Context, oid int64, uid int64) (payload OrderPayload, err error) {
	payload.ID = oid
	payload.UID = uid

	rows, err := store.db.QueryContext(ctx, getOrderInfo, oid, uid)
	if err != nil {
		return
	}
	defer rows.Close()

	var items []OrderItemDetail

	for rows.Next() {
		var i OrderItemDetail
		if err = rows.Scan(&i.PID, &i.Amount); err != nil {
			return
		}
		items = append(items, i)
	}
	if err = rows.Close(); err != nil {
		return
	}

	if err = rows.Err(); err != nil {
		return
	}

	payload.Items = items
	return payload, nil
}

const updateOrderStatus = `
  UPDATE orders 
  SET status = $2
  WHERE id = $1
`

func (store *Store) UpdateOrderStatus(ctx context.Context, id int64, status string) error {
	_, err := store.db.ExecContext(ctx, updateOrderStatus, id, status)
	return err
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
