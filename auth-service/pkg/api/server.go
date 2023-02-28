package api

import (
	"authentication/pkg/db"
	"authentication/pkg/rabbitmq"
	"authentication/pkg/token"
	"authentication/pkg/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	router     *gin.Engine
	store      *db.Store
	tokenMaker token.Maker
	rabbit     *rabbitmq.Rabbit
}

func NewServer(config util.Config, store *db.Store, tokenMaker token.Maker, rabbit *rabbitmq.Rabbit) *Server {

	server := &Server{
		config:     config,
		router:     gin.Default(),
		store:      store,
		tokenMaker: tokenMaker,
		rabbit:     rabbit,
	}

	server.setupRouter()

	return server
}

func (server *Server) Start(port string) error {
	return server.router.Run(port)
}
