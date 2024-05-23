package di

import (
	"github.com/booscaa/order-process-go/microservices/order-service/internal/adapter/http/rest/ordercontroller"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/adapter/http/rest/ordercontroller/database/orderdbrepository"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/adapter/kafka/orderqueuerepository"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/domain"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/usecase/orderusecase"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
)

func ConfigOrderDIController(
	database *sqlx.DB,
	kafkaWriter *kafka.Writer,
) domain.OrderController {
	orderDBRepository := orderdbrepository.New(database)
	orderQueueRepository := orderqueuerepository.New(kafkaWriter)
	orderUseCase := orderusecase.New(orderDBRepository, orderQueueRepository)
	orderController := ordercontroller.New(orderUseCase)

	return orderController
}
