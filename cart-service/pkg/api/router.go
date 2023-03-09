package api

func (server *Server) setupRouter() {

	// handle cors
	server.router.Use(GinMiddleware("http://localhost:3000"))
	v1 := server.router.Group("/v1")
	{
		cart := v1.Group("/cart")
		// check if vailded access token is provided
		cart.Use(authMiddleware(server.config.GRPC_URL))
		{
			cart.GET("", server.ListOwnCarts)
			cart.POST("", server.CreateCart)
			cart.PATCH("", server.UpdateCart)
			cart.DELETE("", server.DeleteCartById)
			cart.DELETE("/batch", server.DeleteBatchCart)
		}
	}
}
