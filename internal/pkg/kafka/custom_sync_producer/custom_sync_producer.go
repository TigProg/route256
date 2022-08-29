package custom_sync_producer

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
)

func New(brokers []string) (*sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = configPkg.KafkaProducerReturnSuccesses

	syncProducer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, err
	}

	log.Debug("kafka producer created")
	return &syncProducer, nil
}
