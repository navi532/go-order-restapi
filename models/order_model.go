package models

import (
	"context"
	"errors"
	"time"

	"gin-order-restapi/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var orderCollection = configs.GetCollection("orders")

type Order struct {
	ID       string `bson:"_id,omitempty"`
	Product  string `bson:"product,omitempty" json:"product,omitempty"`
	Quantity int    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Status   string `bson:"status,omitempty"  json:"status,omitempty" `
}

func (o *Order) Create() (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := orderCollection.InsertOne(ctx, o)
	if err != nil {
		return o, err
	}

	o.ID = (result.InsertedID.(primitive.ObjectID)).String()

	return o, nil
}

func (o *Order) GetAll() ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var orders []*Order

	results, err := orderCollection.Find(ctx, bson.M{})
	if err != nil {
		return orders, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var order Order
		err = results.Decode(&order)

		if err != nil {
			return orders, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func (o *Order) GetOrderById(id string) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order *Order

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, errors.New("invalid id")
	}

	err = orderCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&order)
	if err != nil {
		return order, errors.New("order not found")
	}

	return order, nil
}

func (o *Order) EditOrder(neworder *Order) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order *Order

	objId, err := primitive.ObjectIDFromHex(o.ID)
	if err != nil {
		return order, errors.New("invalid id")
	}

	result, err := orderCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": *neworder})
	if err != nil {
		return order, err
	}

	if result.MatchedCount == 0 {
		return order, errors.New("no order updated")
	}

	return o.GetOrderById(o.ID)
}
