package router

import (
	"sesi10/controllers"
	"sesi10/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.UserRegister)
		userGroup.POST("/login", controllers.UserLogin)
	}

	productGroup := r.Group("/products")
	{
		productGroup.Use(middlewares.Authentication())
		productGroup.POST("/", controllers.CreateProduct)

		productGroup.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
