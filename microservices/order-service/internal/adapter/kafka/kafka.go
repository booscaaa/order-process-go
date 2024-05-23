package kafka

import "github.com/segmentio/kafka-go"

func Initialize(kafkaURL string) *kafka.Writer {
	return &kafka.Writer{
		Addr:      kafka.TCP(kafkaURL),
		Balancer:  &kafka.CRC32Balancer{},
		BatchSize: 1,
	}
}
