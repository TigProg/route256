package rw_repo

import (
	"context"
	"encoding/json"
	"go.opencensus.io/trace"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	kafkaPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka"
)

func New(readRepo repoPkg.Interface, customSyncProducer sarama.SyncProducer, topic string) repoPkg.Interface {
	return &repo{
		readRepo:     readRepo,
		syncProducer: customSyncProducer,
		topic:        topic,
	}
}

type repo struct {
	readRepo     repoPkg.Interface
	syncProducer sarama.SyncProducer
	topic        string
}

func (r *repo) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	return r.readRepo.List(ctx, offset, limit)
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	addMsg := kafkaPkg.AddMessage{Bb: bb}

	newCtx, span := trace.StartSpan(ctx, "Add")
	defer span.End()

	err := r.sendSpecificMessage(newCtx, addMsg)
	if err != nil {
		log.Errorf("rw_repo::Add error %s", err.Error())
		return 0, err
	}
	return 0, nil // TODO
}

func (r *repo) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	return r.readRepo.Get(ctx, id)
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	changeSeatMsg := kafkaPkg.ChangeSeatMessage{
		Id:      id,
		NewSeat: newSeat,
	}

	newCtx, span := trace.StartSpan(ctx, "ChangeSeat")
	defer span.End()

	err := r.sendSpecificMessage(newCtx, changeSeatMsg)
	if err != nil {
		log.Errorf("rw_repo::ChangeSeat error %s", err.Error())
		return err
	}
	return nil // TODO
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	changeDateSeatMsg := kafkaPkg.ChangeDateSeatMessage{
		Id:      id,
		NewDate: newDate,
		NewSeat: newSeat,
	}

	newCtx, span := trace.StartSpan(ctx, "ChangeDateSeat")
	defer span.End()

	err := r.sendSpecificMessage(newCtx, changeDateSeatMsg)
	if err != nil {
		log.Errorf("rw_repo::ChangeDateSeat error %s", err.Error())
		return err
	}
	return nil // TODO
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	deleteMsg := kafkaPkg.DeleteMessage{Id: id}

	newCtx, span := trace.StartSpan(ctx, "Delete")
	defer span.End()

	err := r.sendSpecificMessage(newCtx, deleteMsg)
	if err != nil {
		log.Errorf("rw_repo::Delete error %s", err.Error())
		return err
	}
	return nil // TODO
}

func (r *repo) sendSpecificMessage(ctx context.Context, msg kafkaPkg.SpecificMessage) error {
	_, span := trace.StartSpan(ctx, "sendSpecificMessage")
	defer span.End()

	SpanContextJsObj, err := json.Marshal(span.SpanContext())
	if err != nil {
		return repoPkg.ErrRepoInternal // TODO
	}

	commonMsg := msg.ToCommon()

	jsObj, err := json.Marshal(commonMsg)
	if err != nil {
		return repoPkg.ErrRepoInternal // TODO
	}

	_, _, err = r.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: r.topic,
		Value: sarama.ByteEncoder(jsObj),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte(kafkaPkg.SpanContextHeaderKey),
				Value: SpanContextJsObj,
			},
		},
	})
	if err != nil {
		return repoPkg.ErrRepoInternal // TODO
	}
	return nil
}
