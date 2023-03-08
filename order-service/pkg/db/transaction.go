package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const createOrder = `
  INSERT INTO orders (
    uid
  ) VALUES (
    $1
  )
  RETURNING id
`

const getProductAmountPrice = `
  SELECT amount, price 
  FROM products
  WHERE id = $1
`

const createOrderItem = `
  INSERT INTO order_items (
    oid, pid, amount
  ) VALUES (
    $1, $2, $3
  )
  RETURNING id
`

type CreateOrderItemParam struct {
	PID    int64 `json:"pid" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

type CreateOrderTxParam struct {
	UID   int64
	Items []CreateOrderItemParam
}

func (store *Store) CreateOrderTx(ctx context.Context, args CreateOrderTxParam) (OrderPayload, error) {
	var orderPayload OrderPayload
	orderPayload.UID = args.UID
	orderPayload.Items = []OrderItemDetail{}

	args.Items = reduceCreateOrderTxParam(args.Items)

	err := store.execTx(ctx, func(tx *sql.Tx) error {
		var err error
		row := tx.QueryRowContext(ctx, createOrder, args.UID)
		err = row.Scan(&orderPayload.ID)
		if err != nil {
			return err
		}

		for _, item := range args.Items {
			var remaining int64
			var itemDetail OrderItemDetail
			itemDetail.PID = item.PID
			itemDetail.OID = orderPayload.ID
			itemDetail.Amount = item.Amount

			row := tx.QueryRowContext(ctx, getProductAmountPrice, item.PID)
			err = row.Scan(&remaining, &itemDetail.Price)
			if err != nil {
				return err
			}
			if remaining < item.Amount {
				return fmt.Errorf("Remaining: %v, Request: %v", remaining, item.Amount)
			}

			row = tx.QueryRowContext(ctx, createOrderItem, orderPayload.ID, item.PID, item.Amount)
			err = row.Scan(&itemDetail.ID)
			if err != nil {
				return err
			}
			orderPayload.Items = append(orderPayload.Items, itemDetail)
		}

		return err
	})

	return orderPayload, err
}

func (store *Store) UpdateBatchProductAmountTx(ctx context.Context, products []ProductPayload) error {
	err := store.execTx(ctx, func(tx *sql.Tx) error {
		var err error

		for _, product := range products {
			_, err = store.db.ExecContext(ctx, updateProductAmount, product.ID, product.Amount, time.Now())
			if err != nil {
				return err
			}
		}
		return err
	})

	return err
}
