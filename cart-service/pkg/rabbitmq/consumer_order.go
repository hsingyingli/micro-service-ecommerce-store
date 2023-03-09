package rabbitmq

import (
	"context"
	"time"
)

type BatchCartPayload struct {
	CIDs []int64 `json:"cids"`
	UID  int64
}

func (rabbit *Rabbit) ConsumeDeleteBatchCart(payload BatchCartPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rabbit.store.DeleteBatchCartTx(ctx, payload.CIDs, payload.UID)

	return err
}
