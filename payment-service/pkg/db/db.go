package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = fn(tx)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
