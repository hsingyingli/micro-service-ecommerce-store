package api

import (
	"payment/pkg/db"
	"payment/pkg/rabbitmq"
	"payment/pkg/util"

	"github.com/gin-gonic/gin"
)

// define server struct
type Server struct {
	config util.Config
	router *gin.Engine
	store  *db.Store
	rabbit *rabbitmq.Rabbit
}

func NewServer(config util.Config, store *db.Store, rabbit *rabbitmq.Rabbit) *Server {

	server := &Server{
		config: config,
		router: gin.Default(),
		store:  store,
		rabbit: rabbit,
	}

	server.setupRouter()

	return server
}

func (server *Server) Start(port string) error {
	return server.router.Run(port)
}
