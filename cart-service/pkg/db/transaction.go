package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

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

const deleteBatchCart = `
  DELETE 
  FROM carts 
  WHERE uid = $1 AND id in (%s)
`

func (store *Store) DeleteBatchCartTx(ctx context.Context, order OrderPayload) error {
	err := store.execTx(ctx, func(tx *sql.Tx) error {
		var err error
		cids := []int64{}
		for _, item := range order.Items {
			cids = append(cids, item.CID)
		}

		placeholder := make([]string, len(cids))
		values := make([]interface{}, len(cids))
		for i, id := range cids {
			placeholder[i] = fmt.Sprintf("$%d", i+2)
			values[i] = id
		}

		query := fmt.Sprintf(deleteBatchCart, strings.Join(placeholder, ","))

		_, err = tx.ExecContext(ctx, query, values...)

		return err
	})
	return err
}
