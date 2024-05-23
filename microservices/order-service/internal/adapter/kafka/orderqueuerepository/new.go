package orderqueuerepository

import (
	"context"
	"encoding/json"

	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/domain"
	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/dto"
	"github.com/segmentio/kafka-go"
)

type repository struct {
	kafkaWriter *kafka.Writer
}

// Send implements domain.OrderQueueRepository.
func (repository *repository) Send(ctx context.Context, orderQueueOutput dto.OrderQueueOutput) error {
	orderQueueOutputBytes, err := json.Marshal(&orderQueueOutput)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: orderQueueOutputBytes,
	}

	repository.kafkaWriter.Topic = "OrderCreated"
	err = repository.kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func New(kafkaWriter *kafka.Writer) domain.OrderQueueRepository {
	return &repository{kafkaWriter: kafkaWriter}
}

// func Producer() {
// 	// my-cluster-kafka-bootstrap.kafka:9092
// 	kafkaWriter := getKafkaWriter("kafka:9092", "topic2")

// 	producerHandler(kafkaWriter)
// }
