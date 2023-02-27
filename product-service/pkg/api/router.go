package api

func (server *Server) setupRouter() {

	// handle cors
	server.router.Use(GinMiddleware("http://localhost:3000"))
	v1 := server.router.Group("/v1")
	{

		v1.GET("/product", server.GetProduct)
		v1.GET("/product/all", server.ListProducts)

		product := v1.Group("/auth/product")
		// check if vailded access token is provided
		product.Use(authMiddleware(server.config.GRPC_URL))
		{
			product.POST("", server.CreateProduct)
			product.GET("/all", server.ListOwnProducts)
			product.DELETE("", server.DeleteProductById)
		}
	}

}
