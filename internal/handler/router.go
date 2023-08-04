package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	user.POST("/register", h.createUser)

	user.POST("/login", h.loginUser)

	apiV1.Use(h.authMiddleware())
	admin := apiV1.Group("user/admin")

	admin.POST("/items", h.createItem)
	admin.GET("/items", h.getItems)
	admin.GET("/items/:id", h.getItemByID)
	admin.DELETE("/items/:id", h.deleteItem)
	admin.PUT("/items/:id", h.updateItem)
	return router
}
