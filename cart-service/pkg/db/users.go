package db

import "context"

const createUser = `
  INSERT INTO users (
    id, username, email
  ) VALUES (
    $1, $2, $3
  )
  RETURNING id, username, email, created_at, updated_at
`

type CreateUserParam struct {
	ID       int64
	Username string
	Email    string
}

func (store *Store) CreateUser(ctx context.Context, args CreateUserParam) (User, error) {
	row := store.db.QueryRowContext(ctx, createUser, args.ID, args.Username, args.Email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

const deleteUser = `
  DELETE 
  FROM users 
  WHERE id = $1
`

func (store *Store) DeleteUser(ctx context.Context, id int64) error {
	_, err := store.db.ExecContext(ctx, deleteUser, id)
	return err
}
