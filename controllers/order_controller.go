package controllers

import (
	"context"
	"net/http"
	"time"

	"gin-order-restapi/configs"
	"gin-order-restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	orderCollection = configs.GetCollection("orders")
	validate        = validator.New()
)

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var order models.Order
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&order); err != nil {
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
		if err := validate.Struct(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data":    gin.H{"error_message": err.Error()},
			})
			return
		}

		order.Status = "Processing"

		result, err := orderCollection.InsertOne(ctx, order)
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
			"data":    gin.H{"_id": result.InsertedID},
		})
	}
}

func EditOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderid := c.Param("id")

		var order models.Order
		defer cancel()

		objId, err := primitive.ObjectIDFromHex(orderid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data": gin.H{
					"error_message": "Invalid ID",
				},
			})

			return
		}

		// validate the request body
		if err := c.BindJSON(&order); err != nil {
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
		if err := validate.Struct(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data":    gin.H{"error_message": err.Error()},
			})
			return
		}

		updatedBSONOrder := bson.M{"product": order.Product, "quantity": order.Quantity, "status": order.Status}
		result, err := orderCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatedBSONOrder})
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
		if result.MatchedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data": gin.H{
					"error_message": "No Order with given id",
				},
			})
			return
		}
		var updatedOrder models.Order
		err = orderCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedOrder)
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
			"data": gin.H{
				"order": updatedOrder,
			},
		})
	}
}

func GetSpecificOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderid := c.Param("id")

		var order models.Order
		defer cancel()

		objId, err := primitive.ObjectIDFromHex(orderid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data": gin.H{
					"error_message": "Invalid ID",
				},
			})

			return
		}

		err = orderCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "error",
				"data": gin.H{
					"error_message": "No Order with given id",
				},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data": gin.H{
				"order": order,
			},
		})
	}
}

func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var orders []models.Order
		defer cancel()

		results, err := orderCollection.Find(ctx, bson.M{})
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
		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleOrder models.Order
			err = results.Decode(&singleOrder)

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

			orders = append(orders, singleOrder)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data": gin.H{
				"orders": orders,
			},
		})
	}
}
