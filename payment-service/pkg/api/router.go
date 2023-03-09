package api

func (server *Server) setupRouter() {

	// handle cors
	server.router.Use(GinMiddleware(server.config.ALLOW_ORIGIN))
	v1 := server.router.Group("/v1")
	{
		payment := v1.Group("/payment")
		// check if vailded access token is provided
		payment.Use(authMiddleware(server.config.GRPC_URL))
		{
			payment.POST("", server.FinishPayment)
		}
	}

}
