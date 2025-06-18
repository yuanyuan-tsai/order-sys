package routes

import (
	"order-sys/controller"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	public := r.Group("/api")
	public.GET("/menu", controller.GetMenu)
	// public.POST("/dish", controller.AddMenuItem)
}
