package http

func (s *Server) InitRoutes() {
	v1 := s.router.Group("api/v1")
	items := s.router.Group("api/v2")
	v1.Use(s.handler.UserIdentity)
	items.Use(s.handler.UserWeakIdentity)
	auth := s.router.Group("/auth")
	auth.POST("/sign-in", s.handler.SignIn)
	auth.POST("/sign-up", s.handler.SignUp)

	v1.PUT("/user", s.handler.UpdateUser)
	v1.GET("/user", s.handler.GetUserById)
	v1.PUT("/user/change-password", s.handler.UpdatePassword)
	v1.DELETE("/user", s.handler.DeleteUser)
	v1.GET("/user/favorites", s.handler.GetUserFavoriteProducts)
	v1.GET("/user/purchases", s.handler.GetUserProducts)
	v1.POST("/user/deposit", s.handler.DepositBalance)
	v1.POST("/products/:product_id/like", s.handler.LikeAction)
	v1.POST("/products/:product_id/purchase", s.handler.BuyProduct)

	items.GET("/products", s.handler.GetProducts)
	items.POST("/products", s.handler.CreateProduct)
	items.DELETE("/products/:id", s.handler.DeleteProduct)
	items.PUT("/products/:id", s.handler.UpdateProduct)
	items.GET("/products/:id", s.handler.GetProductById)
	items.GET("/categories", s.handler.GetCategories)
	items.GET("/categories/:category_id/products", s.handler.GetCategoryProducts)
}
