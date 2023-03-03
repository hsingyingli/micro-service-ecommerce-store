package api

func (server *Server) setupRouter() {

	// handle cors
	server.router.Use(GinMiddleware("http://localhost:3000"))
	v1 := server.router.Group("/v1")
	{
		order := v1.Group("/order")
		// check if vailded access token is provided
		order.Use(authMiddleware(server.config.GRPC_URL))
		{
			order.GET("", server.ListOrders)
			order.POST("", server.CreateOrder)
			order.DELETE("", server.DeleteOrder)
		}
	}
}
