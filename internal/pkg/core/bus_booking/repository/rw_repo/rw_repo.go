package rw_repo

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
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
	err := r.sendSpecificMessage(addMsg)
	if err != nil {
		return 0, err // TODO
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
	err := r.sendSpecificMessage(changeSeatMsg)
	if err != nil {
		return err // TODO
	}
	return nil // TODO
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	changeDateSeatMsg := kafkaPkg.ChangeDateSeatMessage{
		Id:      id,
		NewDate: newDate,
		NewSeat: newSeat,
	}
	err := r.sendSpecificMessage(changeDateSeatMsg)
	if err != nil {
		return err // TODO
	}
	return nil // TODO
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	deleteMsg := kafkaPkg.DeleteMessage{Id: id}
	err := r.sendSpecificMessage(deleteMsg)
	if err != nil {
		return err // TODO
	}
	return nil // TODO
}

func (r *repo) sendSpecificMessage(msg kafkaPkg.SpecificMessage) error {
	commonMsg := msg.ToCommon()

	jsObj, err := json.Marshal(commonMsg)
	if err != nil {
		return repoPkg.ErrRepoInternal // TODO
	}

	_, _, err = r.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: r.topic,
		Value: sarama.ByteEncoder(jsObj),
	})
	if err != nil {
		return repoPkg.ErrRepoInternal // TODO
	}
	return nil
}
