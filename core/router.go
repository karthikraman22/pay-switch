package core

import (
	"context"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

type Router struct {
	kafkaUrl  string
	producers map[string]*kafka.Writer
}

func NewRouter(kafkaUrl string) *Router {
	return &Router{kafkaUrl: kafkaUrl, producers: make(map[string]*kafka.Writer)}
}

func (r *Router) AddRoute(topic string) {
	r.producers[topic] = getKafkaWriter(r.kafkaUrl, topic)
}

func (r *Router) Route(payload []byte, headers map[string]string, ctx context.Context) (string, error) {
	platformId := uuid.NewString()
	kHeaders := make([]protocol.Header, len(headers))
	for k, v := range headers {
		hdr := protocol.Header{Key: k, Value: []byte(v)}
		kHeaders = append(kHeaders, hdr)
	}
	msg := kafka.Message{
		Key:     []byte(platformId),
		Headers: kHeaders,
		Value:   payload,
	}
	err := r.producers["payment.initiate"].WriteMessages(ctx, msg)
	return platformId, err
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
