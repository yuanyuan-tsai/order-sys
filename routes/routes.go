package routes

import (
	"order-sys/controller"
	"order-sys/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.GET("/menu", controller.GetMenu)
	}
	{
		user_public := public.Group("/user")
		user_public.POST("/register", controller.Register)
		user_public.POST("/login", controller.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		user_protected := protected.Group("/user")
		user_protected.POST("/order", controller.CreateOrder)
		user_protected.GET("/orders", controller.GetUserAllOrders)
		user_protected.GET("/orders/:id", controller.GetUserOrderByID)
	}
	{
		chef_protected := protected.Group("/chef")
		chef_protected.GET("/orders", controller.GetPendingOrders)
		chef_protected.POST("/orders/:id/:itemId", controller.UpdateDishStatus)
	}
	{
		admin_protected := protected.Group("/admin")
		admin_protected.POST("/dish", controller.AddMenuItem)
	}
}
