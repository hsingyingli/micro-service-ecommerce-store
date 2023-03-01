package db

import (
	"context"
	"time"
)

const createUser = `
INSERT INTO users ( 
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING id, username, email, password, created_at, updated_at
`

type CreateUserParam struct {
	Username string
	Email    string
	Password string
}

func (store *Store) CreateUser(ctx context.Context, args CreateUserParam) (User, error) {
	row := store.db.QueryRowContext(ctx, createUser, args.Username, args.Email, args.Password)
	var user User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
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

const getUserByEmail = `
  SELECT *
  FROM users 
  WHERE email = $1 LIMIT 1
  `

func (store *Store) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := store.db.QueryRowContext(ctx, getUserByEmail, email)
	var user User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

const getUserById = `
  SELECT *
  FROM users
  WHERE id = $1 LIMIT 1
  `

func (store *Store) GetUserById(ctx context.Context, id int64) (User, error) {
	row := store.db.QueryRowContext(ctx, getUserById, id)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

const updateUserInfo = `
  UPDATE users 
  SET username = $2,
      email = $3,
      updated_at = $4
  WHERE id = $1
  RETURNING id, username, email, password, created_at, updated_at
`

type UpdateUserInfoParam struct {
	ID       int64
	Username string
	Email    string
}

func (store *Store) UpdateUserInfo(ctx context.Context, args UpdateUserInfoParam) (User, error) {
	row := store.db.QueryRowContext(ctx, updateUserInfo, args.ID, args.Username, args.Email, time.Now())
	var user User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

const updateUserPassword = `
  UPDATE users 
  SET password = $2,
      updated_at = $3
  WHERE id = $1
  RETURNING id, username, email, password, created_at, updated_at
`

type UpdateUserPasswordParam struct {
	ID       int64
	Password string
}

func (store *Store) UpdateUserPassword(ctx context.Context, args UpdateUserPasswordParam) (User, error) {
	row := store.db.QueryRowContext(ctx, updateUserPassword, args.ID, args.Password, time.Now())
	var user User

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
