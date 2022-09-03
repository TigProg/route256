package custom_consumer

import (
	"context"
	"encoding/json"
	"errors"
	"go.opencensus.io/trace"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	kafkaPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka"
	metricPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/metrics"
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

	log.Debug("kafka consumer created")
	return &Consumer{
		client: client,
		repo:   repo,
		metrics: metricManager{
			success: metricPkg.NewMetric("custom_consumer::success"),
			fail:    metricPkg.NewMetric("custom_consumer::fail"),
			income:  metricPkg.NewMetric("custom_consumer::income"),
		},
	}, nil
}

// TODO change structure to remove copy-paste
type metricManager struct {
	success *metricPkg.Metric
	fail    *metricPkg.Metric
	income  *metricPkg.Metric
}

type Consumer struct {
	client  sarama.ConsumerGroup
	repo    repoPkg.Interface
	metrics metricManager
}

func (c *Consumer) GetMetrics() []*metricPkg.Metric {
	return []*metricPkg.Metric{
		c.metrics.success,
		c.metrics.fail,
		c.metrics.income,
	}
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
			log.Info("consumer session done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Info("consumer data channel closed")
				return nil
			} else {
				c.metrics.income.Increment()
				err := c.handle(msg)
				if err == nil {
					c.metrics.success.Increment()
				} else {
					c.metrics.fail.Increment()
					log.Errorf("error on handle %v: %v", msg.Value, err)
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
			log.Errorf("consumer failed on consume: %v", err)
			time.Sleep(consumerSleep)
		}
	}
}

func (c *Consumer) handle(msg *sarama.ConsumerMessage) error {
	value := msg.Value
	headers := msg.Headers

	var spanContext trace.SpanContext
	for _, header := range headers {
		if string(header.Key[:]) == kafkaPkg.SpanContextHeaderKey {
			if err := json.Unmarshal(header.Value, &spanContext); err != nil {
				log.Error("broken context")
			}
			break
		}
	}

	commonMsg := kafkaPkg.CommonMessage{}
	err := json.Unmarshal(value, &commonMsg)
	if err != nil {
		log.Errorf("consumer failed on unmarshal: %v", err)

		return err // TODO custom error
	}

	specificMsg, err := commonMsg.ToSpecificMessage()
	if err != nil {
		return err // TODO custom error
	}

	newCtx, span := trace.StartSpanWithRemoteParent(
		context.Background(), "handle", spanContext,
	)
	defer span.End()

	// TODO add single handler
	switch commonMsg.Key {
	case kafkaPkg.AddKey:
		addMsg := specificMsg.(kafkaPkg.AddMessage)
		_, err := c.repo.Add(newCtx, addMsg.Bb)
		return err
	case kafkaPkg.ChangeSeatKey:
		changeSeatMsg := specificMsg.(kafkaPkg.ChangeSeatMessage)
		return c.repo.ChangeSeat(
			newCtx,
			changeSeatMsg.Id,
			changeSeatMsg.NewSeat,
		)
	case kafkaPkg.ChangeDateSeatKey:
		changeDateSeatMsg := specificMsg.(kafkaPkg.ChangeDateSeatMessage)
		return c.repo.ChangeDateSeat(
			newCtx,
			changeDateSeatMsg.Id,
			changeDateSeatMsg.NewDate,
			changeDateSeatMsg.NewSeat,
		)
	case kafkaPkg.DeleteKey:
		deleteMsg := specificMsg.(kafkaPkg.DeleteMessage)
		return c.repo.Delete(
			newCtx,
			deleteMsg.Id,
		)
	default:
		return errors.New("unreachable code") // TODO
	}
}
