package db

import "context"

const getOrderInfo = `
  SELECT oi.id, oi.pid, oi.amount, p.price, p.title, p.image_name
  FROM order_items as oi
  JOIN products as p ON oi.pid = p.id
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
		if err = rows.Scan(&i.ID, &i.PID, &i.Amount, &i.Price, &i.Title, &i.ImageName); err != nil {
			return
		}
		i.OID = oid
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

const listOrder = `
  SELECT id 
  FROM orders 
  WHERE uid = $1 AND status = 'WAIT'
`

func (store *Store) ListOrderInfo(ctx context.Context, uid int64) ([]OrderPayload, error) {
	rows, err := store.db.QueryContext(ctx, listOrder, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orderId []int64

	for rows.Next() {
		var i int64
		if err := rows.Scan(&i); err != nil {
			return nil, err
		}
		orderId = append(orderId, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var payloadList []OrderPayload
	for _, i := range orderId {
		payload, err := store.GetOrderInfo(ctx, i, uid)
		if err != nil {
			return nil, err
		}
		payloadList = append(payloadList, payload)
	}
	return payloadList, nil
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
