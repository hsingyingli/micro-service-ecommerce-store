package rabbitmq

import (
	"context"
	"log"
	"order/pkg/db"
	"time"
)

type UserPayload struct {
	ID       int64
	Username string
	Email    string
}

func (rabbit *Rabbit) ConsumeCreateUser(user UserPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rabbit.store.CreateUser(ctx, db.CreateUserParam{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
	return err
}

func (rabbit *Rabbit) ConsumeUpdateUser(user UserPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rabbit.store.UpdateUser(ctx, db.UpdateUserParam{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	log.Println(err)
	return err
}

func (rabbit *Rabbit) ConsumeDeleteUser(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rabbit.store.DeleteUser(ctx, id)
	return err
}
