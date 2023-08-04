package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mereiamangeldin/One-lab-Homework-1/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	user.POST("/register", h.createUser)

	user.POST("/login", h.loginUser)

	apiV1.Use(h.authMiddleware())
	admin := apiV1.Group("user/admin")

	admin.POST("/items", h.createItem)
	admin.DELETE("/items/:id", h.deleteItem)
	admin.PUT("/items/:id", h.updateItem)

	apiV1.GET("/items", h.getItems)
	apiV1.GET("/items/:id", h.getItemByID)
	apiV1.GET("/items/:id/purchase", h.buyItem)
	apiV1.GET("my-items", h.userItems)
	return router
}
