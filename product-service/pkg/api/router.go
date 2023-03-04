package api

func (server *Server) setupRouter() {

	// handle cors
	server.router.Use(GinMiddleware("http://localhost:3000"))
	v1 := server.router.Group("/v1")
	{

		product := v1.Group("/product")
		product.GET("", server.GetProduct)
		product.GET("/all", server.ListProducts)

		auth := product.Group("/auth")
		// check if vailded access token is provided
		auth.Use(authMiddleware(server.config.GRPC_URL))
		{
			auth.POST("", server.CreateProduct)
			auth.GET("/all", server.ListOwnProducts)
			auth.DELETE("", server.DeleteProductById)
		}
	}

}
