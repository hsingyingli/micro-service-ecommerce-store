package api

import (
	"authentication/pkg/db"
	"authentication/pkg/rabbitmq"
	"authentication/pkg/token"
	"authentication/pkg/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserInfoRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (server *Server) UpdateUserInfo(ctx *gin.Context) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var req UpdateUserInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := server.store.UpdateUserInfo(ctx, db.UpdateUserInfoParam{
		ID:       payload.UID,
		Username: req.Username,
		Email:    req.Email,
	})

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishUser(ctx, "user.update", rabbitmq.UserPayload{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rsp := UserResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
	}
	ctx.JSON(http.StatusOK, rsp)

}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

func (server *Server) UpdateUserPassword(ctx *gin.Context) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var req UpdateUserPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser, err := server.store.UpdateUserPassword(ctx, db.UpdateUserPasswordParam{
		ID:       payload.UID,
		Password: hashedPassword,
	})

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rsp := UserResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) DeleteUser(ctx *gin.Context) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	err := server.store.DeleteUser(ctx, payload.UID)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = server.rabbit.PublishUser(ctx, "user.delete", rabbitmq.UserPayload{
		ID: payload.UID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}
