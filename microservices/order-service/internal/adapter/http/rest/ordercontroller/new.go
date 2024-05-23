package ordercontroller

import (
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/domain"
)

type controller struct {
	usecase domain.OrderUseCase
}

func New(usecase domain.OrderUseCase) domain.OrderController {
	return &controller{
		usecase: usecase,
	}
}
