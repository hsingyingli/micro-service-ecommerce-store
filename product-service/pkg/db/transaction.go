package db

import (
	"database/sql"
	"errors"

	"golang.org/x/net/context"
)

type OrderPayload struct {
	ID     int64
	PID    int64
	UID    int64 // product buyer
	Amount int64
	Price  int64
}

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

type ProductPayload struct {
	ID        int64
	UID       int64
	Title     string
	Price     int64
	Amount    int64
	ImageName string
}

func (store *Store) UpdateInventoryStatuTx(ctx context.Context, pid int64, a int64, n int64) (ProductPayload, error) {
	var product ProductPayload

	err := store.execTx(ctx, func(tx *sql.Tx) error {
		var err error
		var amount int64
		var numUnPaid int64

		row := tx.QueryRowContext(ctx, getProductAmount, pid)
		err = row.Scan(&amount, &numUnPaid)
		if err != nil {
			return err
		}

		amount += a
		numUnPaid += n

		if amount < 0 || numUnPaid < 0 {
			return errors.New("Inventory Status went wrong!!")
		}

		row = tx.QueryRowContext(ctx, updateProductAmount, pid, amount, numUnPaid)
		err = row.Scan(
			&product.ID,
			&product.UID,
			&product.Amount,
			&product.Price,
		)

		return err
	})
	return product, err
}
