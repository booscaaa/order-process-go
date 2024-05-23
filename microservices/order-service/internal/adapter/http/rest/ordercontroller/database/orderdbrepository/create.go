package orderdbrepository

import (
	"context"

	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/domain"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/dto"
)

// Create implements domain.OrderDBRepository.
func (repository *repository) Create(ctx context.Context, orderInput dto.OrderInput) (*domain.Order, error) {
	query := `INSERT INTO order (product_id, qtt) VALUES ($1, $2) returning *;`

	orderCreated := domain.Order{}

	err := repository.database.QueryRowx(
		query,
		orderInput.ProductID,
		orderInput.QTT,
	).StructScan(&orderCreated)

	if err != nil {
		return nil, err
	}

	return &orderCreated, nil
}
