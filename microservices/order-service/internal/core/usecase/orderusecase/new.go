package orderusecase

import (
	"context"

	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/domain"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/dto"
)

type usecase struct {
	orderDBRepostory     domain.OrderDBRepository
	orderQueueRepository domain.OrderQueueRepository
}

// Create implements domain.OrderUseCase.
func (usecase *usecase) Create(ctx context.Context, orderInput dto.OrderInput) (*domain.Order, error) {
	orderCreated, err := usecase.orderDBRepostory.Create(ctx, orderInput)
	if err != nil {
		return nil, err
	}

	err = usecase.orderQueueRepository.Send(ctx, dto.OrderQueueOutput{
		ProductID: orderCreated.ProductID,
		QTT:       orderCreated.QTT,
	})
	if err != nil {
		return nil, err
	}

	return orderCreated, nil
}

func New(
	orderDBRepostory domain.OrderDBRepository,
	orderQueueRepository domain.OrderQueueRepository,
) domain.OrderUseCase {
	return &usecase{
		orderDBRepostory:     orderDBRepostory,
		orderQueueRepository: orderQueueRepository,
	}
}
