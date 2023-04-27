package requests

type CreateOrderRequest struct {
	Product  string `bson:"product" json:"product" validate:"required"`
	Quantity int    `bson:"quantity" json:"quantity" validate:"required,gte=1"`
}

type EditOrderRequest struct {
	Product  string `json:"product,omitempty"`
	Quantity int    `json:"quantity" validate:"omitempty,gte=1"`
	Status   string `json:"status,omitempty"`
}

type GetOrderRequest struct{}
