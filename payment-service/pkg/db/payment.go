package db

import (
	"context"
	"database/sql"
)

const createPayment = `
  INSERT INTO payments (
    uid, oid
  ) VALUES (
    $1, $2
  )
  `

func (store *Store) CreatePayment(ctx context.Context, uid int64, oid int64) error {
	err := store.execTx(ctx, func(tx *sql.Tx) error {
		var err error
		_, err = tx.ExecContext(ctx, createPayment, uid, oid)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, deleteOrder, oid, uid)

		return err
	})
	return err
}
