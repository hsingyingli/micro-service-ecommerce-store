package api

import (
	"cart/pkg/db"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) ListOwnCarts(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)

	carts, err := server.store.ListCarts(ctx, user.Id)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, carts)
}

type CreateCartRequest struct {
	PID    int64 `json:"pid" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

func (server *Server) CreateCart(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)

	var req CreateCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	remaining, err := server.store.GetProductAmountById(ctx, req.PID)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if remaining < req.Amount {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "shortage of stock"})
		return
	}

	cart, err := server.store.CreateCart(ctx, db.CreateCartParam{
		UID:    user.Id,
		PID:    req.PID,
		Amount: req.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cart)
}

type UpdateCartRequest struct {
	PID    int64 `json:"pid" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

func (server *Server) UpdateCart(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	var req UpdateCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	remaining, err := server.store.GetProductAmountById(ctx, req.PID)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if remaining < req.Amount {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "shortage of stock"})
		return
	}

	cart, err := server.store.UpdateCart(ctx, req.PID, user.Id, req.Amount)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, cart)
}

func (server *Server) DeleteCartById(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	i := ctx.Query("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.store.DeleteCart(ctx, int64(id), user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
