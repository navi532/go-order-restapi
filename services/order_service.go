package services

import (
	"encoding/json"

	"gin-order-restapi/controllers/requests"
	"gin-order-restapi/models"
)

func CreateOrder(request *requests.CreateOrderRequest) (*models.Order, error) {
	var order *models.Order

	order = &models.Order{
		Product:  request.Product,
		Status:   "Processing",
		Quantity: request.Quantity,
	}

	order, err := order.Create()

	return order, err
}

func GetAllOrders(request *requests.GetOrderRequest) ([]*models.Order, error) {
	order := &models.Order{}
	orders, err := order.GetAll()

	return orders, err
}

func GetOrderById(request *requests.GetOrderRequest, id string) (*models.Order, error) {
	order := &models.Order{}
	orders, err := order.GetOrderById(id)

	return orders, err
}

func EditOrder(request *requests.EditOrderRequest, id string) (*models.Order, error) {
	order := &models.Order{}

	order, err := order.GetOrderById(id)
	if err != nil {
		return order, nil
	}

	// convert request to neworder
	neworder := &models.Order{}
	val, err := json.Marshal(*request)
	if err != nil {
		return order, nil
	}
	err = json.Unmarshal(val, neworder)
	if err != nil {
		return order, nil
	}

	order, err = order.EditOrder(neworder)

	return order, err
}
