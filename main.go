package main

import (
	"github.com/navi532/gin-order-restapi/configs"
	"github.com/navi532/gin-order-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.OrderRoute(router)

	router.Run("localhost:8080")
}
