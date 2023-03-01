package rabbitmq

import (
	"cart/pkg/db"
	"context"
	"log"
	"time"
)

type UserPayload struct {
	ID       int64
	Username string
	Email    string
}

func (consumer *Consumer) CreateUser(user UserPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := consumer.store.CreateUser(ctx, db.CreateUserParam{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
	return err
}

func (consumer *Consumer) UpdateUser(user UserPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println(user)

	_, err := consumer.store.UpdateUser(ctx, db.UpdateUserParam{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	log.Println(err)
	return err
}

func (consumer *Consumer) DeleteUser(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := consumer.store.DeleteUser(ctx, id)
	return err
}
