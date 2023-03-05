package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RenewAccessTokenResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       int64  `json:"id"`
}

// RenewAccessToken return access token if a vailded refresh token is provided
func (server *Server) RenewAccessToken(ctx *gin.Context) {

	// check if refresh token is provided or not
	refreshToken, err := ctx.Cookie("ecommerce-store-refresh-token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	log.Println("receive token")

	// verify token
	payload, err := server.tokenMaker.VerifyToken(refreshToken)

	if err != nil {
		// invaild token or has expired
		ctx.SetCookie("ecommerce-store-refresh-token", "", -1, "/", "localhost", false, true)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// get user data
	user, err := server.store.GetUserById(ctx, payload.UID)

	if err != nil {
		ctx.SetCookie("ecommerce-store-refresh-token", "", -1, "/", "localhost", false, true)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	// create new access token and return
	accessToken, err := server.tokenMaker.CreateToken(user, server.config.ACCESS_TOKEN_DURATION)

	rsp := RenewAccessTokenResponse{
		Username: user.Username,
		Email:    user.Email,
		Id:       user.ID,
	}

	maxage := server.config.ACCESS_TOKEN_DURATION.Microseconds()
	ctx.SetCookie("ecommerce-store-access-token", accessToken, int(maxage), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, rsp)
}
