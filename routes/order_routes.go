package routes

import (
	"gin-order-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.GET("/order", controllers.GetAllOrders())
	router.GET("/order/:id", controllers.GetSpecificOrder())
	router.PUT("/order/:id", controllers.EditOrder())
	router.POST("/order", controllers.CreateOrder())
}
