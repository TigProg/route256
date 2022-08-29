package custom_consumer

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/Shopify/sarama"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	kafkaPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka"
)

// TODO change type to interface

func New(brokers []string, repo repoPkg.Interface, groupId string) (*Consumer, error) {
	cfg := sarama.NewConfig()
	{
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
		cfg.Consumer.Return.Errors = configPkg.KafkaConsumerReturnErrors
	}

	client, err := sarama.NewConsumerGroup(brokers, groupId, cfg)
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
				err := c.handle(msg.Value)
				if err != nil {
					log.Panic(err)
				}
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

func (c *Consumer) handle(value []byte) error {
	commonMsg := kafkaPkg.CommonMessage{}
	err := json.Unmarshal(value, &commonMsg)
	if err != nil {
		return err // TODO custom error
	}

	specificMsg, err := commonMsg.ToSpecificMessage()
	if err != nil {
		return err // TODO custom error
	}

	ctx := context.Background() // TODO

	// TODO add single handler
	switch commonMsg.Key {
	case kafkaPkg.AddKey:
		addMsg := specificMsg.(kafkaPkg.AddMessage)
		_, err := c.repo.Add(ctx, addMsg.Bb)
		return err
	case kafkaPkg.ChangeSeatKey:
		changeSeatMsg := specificMsg.(kafkaPkg.ChangeSeatMessage)
		return c.repo.ChangeSeat(
			ctx,
			changeSeatMsg.Id,
			changeSeatMsg.NewSeat,
		)
	case kafkaPkg.ChangeDateSeatKey:
		changeDateSeatMsg := specificMsg.(kafkaPkg.ChangeDateSeatMessage)
		return c.repo.ChangeDateSeat(
			ctx,
			changeDateSeatMsg.Id,
			changeDateSeatMsg.NewDate,
			changeDateSeatMsg.NewSeat,
		)
	case kafkaPkg.DeleteKey:
		deleteMsg := specificMsg.(kafkaPkg.DeleteMessage)
		return c.repo.Delete(
			ctx,
			deleteMsg.Id,
		)
	default:
		return errors.New("unreachable code") // TODO
	}
}
