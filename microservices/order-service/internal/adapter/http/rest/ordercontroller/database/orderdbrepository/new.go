package orderdbrepository

import (
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) domain.OrderDBRepository {
	return &repository{database: database}
}
