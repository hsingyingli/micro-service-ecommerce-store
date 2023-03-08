package db

import (
	"database/sql"
	"errors"

	"golang.org/x/net/context"
)

const getProductAmount = `
  SELECT amount, num_unpaid
  FROM products
  WHERE id = $1
`

const updateProductAmount = `
  UPDATE products 
  SET amount = $2,
      num_unpaid = $3
  WHERE id = $1
  RETURNING id, uid, amount, price
`



func (store *Store) UpdateInventoryStatuTx(ctx context.Context, order OrderPayload, a int64, n int64) ([]ProductPayload, error) {
	var products []ProductPayload
	err := store.execTx(ctx, func(tx *sql.Tx) error {
		var err error

		for _, item := range order.Items {
			var product ProductPayload
			var amount int64
			var numUnPaid int64
			row := tx.QueryRowContext(ctx, getProductAmount, item.PID)
			err = row.Scan(&amount, &numUnPaid)
			if err != nil {
				return err
			}
			amount += a * item.Amount
			numUnPaid += n * item.Amount
			if amount < 0 || numUnPaid < 0 {
				return errors.New("Inventory Status went wrong!!")
			}
			row = tx.QueryRowContext(ctx, updateProductAmount, item.PID, amount, numUnPaid)
			err = row.Scan(
				&product.ID,
				&product.UID,
				&product.Amount,
				&product.Price,
			)
			if err != nil {
				return err
			}

			products = append(products, product)
		}
		return err
	})

	return products, err
}
