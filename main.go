package main

import (
	"gin-order-restapi/configs"
	"gin-order-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.OrderRoute(router)

	router.Run("localhost:8080")
}
