package redis_wrapper

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	"strconv"
	"time"
)

type Interface interface {
	GetById(id uint) (*models.BusBooking, error)
	SetById(id uint, bb models.BusBooking) error
	DisableById(id uint) error
}

var ErrRedisMiss = errors.New("redis miss")

func New(address string, db int, password string, expiration time.Duration) Interface {
	return &redisWrapper{
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			DB:       db,
			Password: password,
		}),
		expiration: expiration,
	}
}

type redisWrapper struct {
	client     *redis.Client
	expiration time.Duration
}

func (rw *redisWrapper) GetById(id uint) (*models.BusBooking, error) {
	result := rw.client.Get(strconv.Itoa(int(id)))
	if result.Err() != nil {
		log.Errorf("redisWrapper::GetById %s", result.Err().Error())
		return nil, result.Err()
	}

	var bb models.BusBooking
	err := result.Scan(&bb)
	if err != nil {
		log.Infof("redisWrapper::GetById miss")
		return nil, ErrRedisMiss
	}
	return &bb, nil
}

func (rw *redisWrapper) SetById(id uint, bb models.BusBooking) error {
	result := rw.client.Set(strconv.Itoa(int(id)), bb, rw.expiration)
	if result.Err() != nil {
		log.Errorf("redisWrapper::SetById %s", result.Err().Error())
		return result.Err()
	}
	return nil
}

func (rw *redisWrapper) DisableById(id uint) error {
	result := rw.client.Del(strconv.Itoa(int(id)))
	if result.Err() != nil {
		log.Errorf("redisWrapper::DisableById %s", result.Err().Error())
		return result.Err()
	}
	return nil
}
