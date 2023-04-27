package controllers

import (
	"net/http"

	"gin-order-restapi/configs"
	"gin-order-restapi/controllers/requests"
	"gin-order-restapi/services"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

var (
	validate        = validator.New()
	orderCollection = configs.GetCollection("orders")
)

func CreateOrder(c *gin.Context) {
	request := &requests.CreateOrderRequest{}

	// validate the request body
	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error",
			"data": gin.H{
				"error_message": err.Error(),
			},
		})
		return
	}
	// use the validator library to validate required fields
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error",
			"data":    gin.H{"error_message": err.Error()},
		})
		return
	}

	result, err := services.CreateOrder(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error",
			"data":    gin.H{"error_message": err.Error()},
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "success",
		"data":    result,
	})
}

func EditOrder(c *gin.Context) {
	id := c.Param("id")

	request := &requests.EditOrderRequest{}

	// validate the request body
	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error",
			"data": gin.H{
				"error_message": err.Error(),
			},
		})
		return
	}

	// use the validator library to validate required fields
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error",
			"data":    gin.H{"error_message": err.Error()},
		})
		return
	}
	result, err := services.EditOrder(request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error",
			"data":    gin.H{"error_message": err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    result,
	})
}

func GetSpecificOrder(c *gin.Context) {
	id := c.Param("id")

	request := &requests.GetOrderRequest{}

	result, err := services.GetOrderById(request, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
				"data": gin.H{
					"error_message": err.Error(),
				},
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    result,
	})
}

func GetAllOrders(c *gin.Context) {
	request := &requests.GetOrderRequest{}

	result, err := services.GetAllOrders(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error",
			"data":    gin.H{"error_message": err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data": gin.H{
			"count":  len(result),
			"orders": result,
		},
	})
}
