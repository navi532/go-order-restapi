package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID `bson:"_id"`
	Product  string             `bson:"product,omitempty" json:"product,omitempty" validate:"required"`
	Quantity int                `bson:"quantity,omitempty" json:"quantity,omitempty" validate:"required,gte=1"`
	Status   string             `bson:"status,omitempty"  json:"status,omitempty" `
}
