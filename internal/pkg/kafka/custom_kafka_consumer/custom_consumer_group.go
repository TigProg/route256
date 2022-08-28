package custom_kafka_consumer

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

func New(brokers []string, offsetInitial int64, isReturnErrors bool, repo repoPkg.Interface) (sarama.ConsumerGroupHandler, error) {
	cfg := sarama.NewConfig()
	{
		cfg.Consumer.Offsets.Initial = offsetInitial
		cfg.Consumer.Return.Errors = isReturnErrors
	}

	client, err := sarama.NewConsumerGroup(brokers, configPkg.KafkaGroupId1, cfg)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		client: client,
		repo:   repo,
	}, nil
}

type Consumer struct {
	client sarama.ConsumerGroup
	repo   repoPkg.Interface
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case <-session.Context().Done():
			log.Print("Done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Print("data channel closed")
				return nil
			} else {
				log.Printf("%v", msg.Value)
				// TODO
				session.MarkMessage(msg, "")
			}
		}
	}
}

func (c *Consumer) Run(ctx context.Context, topics []string, consumerSleep time.Duration) {
	for {
		err := c.client.Consume(ctx, topics, c)
		if err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(consumerSleep)
		}
	}
}
