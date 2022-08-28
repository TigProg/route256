package custom_sync_producer

import "github.com/Shopify/sarama"

func New(brokers []string, isReturnSuccesses bool) (*sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = isReturnSuccesses

	syncProducer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, err
	}
	return &syncProducer, nil
}
