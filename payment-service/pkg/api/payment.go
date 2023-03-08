package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FinishPaymentRequest struct {
	OID int64 `json:"oid" binding:"required"`
}

func (server *Server) FinishPayment(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	var req FinishPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := server.store.GetOrderById(ctx, req.OID, user.Id)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = server.store.CreatePayment(ctx, user.Id, req.OID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishPayment(ctx, "payment.success", order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}
