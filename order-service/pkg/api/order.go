package api

import (
	"net/http"
	"order/pkg/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Items []db.CreateOrderItemParam `json:"items" binding:"required"`
}

func (server *Server) CreateOrder(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	var req CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderPayload, err := server.store.CreateOrderTx(ctx, db.CreateOrderTxParam{
		UID:   user.Id,
		Items: req.Items,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishOrder(ctx, "order.create", orderPayload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderPayload)
}

func (server *Server) DeleteOrder(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	i := ctx.Query("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderPayload, err := server.store.GetOrderInfo(ctx, int64(id), user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.store.DeleteOrder(ctx, int64(id), user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishOrder(ctx, "order.delete", orderPayload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
