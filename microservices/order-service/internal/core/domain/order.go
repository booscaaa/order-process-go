package domain

import (
	"context"
	"net/http"

	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/dto"
)

type Order struct {
	ID        string
	ProductID string
	QTT       float64
	Price     float64
}

type OrderQueueRepository interface {
	Send(context.Context, dto.OrderQueueOutput) error
}

type OrderDBRepository interface {
	Create(context.Context, dto.OrderInput) (*Order, error)
}

type OrderUseCase interface {
	Create(context.Context, dto.OrderInput) (*Order, error)
}

type OrderController interface {
	Create(http.ResponseWriter, *http.Request)
}

func (order *Order) CalculatePrice(productPrice float64) {
	order.Price = order.QTT * productPrice
}
