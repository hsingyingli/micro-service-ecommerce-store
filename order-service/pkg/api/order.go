package api

import (
	"database/sql"
	"net/http"
	"order/pkg/db"
	"order/pkg/rabbitmq"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) ListOrders(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	l := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(l)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	o := ctx.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(o)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orders, err := server.store.ListOrders(ctx, user.Id, int64(limit), int64(offset))
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

type CreateOrderRequest struct {
	PID    int64 `json:"pid" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

func (server *Server) CreateOrder(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	var req CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := server.store.CreateOrder(ctx, db.CreateOrderParam{
		UID:    user.Id,
		PID:    req.PID,
		Amount: req.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	price, err := server.store.GetProductPriceById(ctx, order.PID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishOrder(ctx, "order.create", rabbitmq.OrderPayload{
		ID:     order.ID,
		PID:    order.PID,
		UID:    user.Id,
		Amount: order.Amount,
		Price:  price,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (server *Server) DeleteOrder(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	i := ctx.Query("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pid, amount, err := server.store.GetOrderAmount(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.store.DeleteOrder(ctx, int64(id), user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishOrder(ctx, "order.delete", rabbitmq.OrderPayload{
		PID:    pid,
		Amount: amount,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
