package config

import "time"

const (
	TelegramBotApiDebug   = true
	TelegramBotApiTimeout = 60

	GRPCClientTarget  = ":8081"
	GRPCServerAddress = ":8081"

	RepoGRPCServerAddress = ":5999"

	RESTServerAddress = ":8088"

	LocalCachePoolSize = 10

	PosgtreSQLHost     = "localhost"
	PosgtreSQLPort     = 5432
	PosgtreSQLUser     = "user"
	PosgtreSQLPassword = "password"
	PosgtreSQLDBname   = "bus_booking"

	PosgtreSQLMaxConnIdleTime = time.Minute
	PosgtreSQLMaxConnLifetime = time.Hour
	PosgtreSQLMinConns        = 2
	PosgtreSQLMaxConns        = 4

	KafkaBrokers                 = "localhost:9095,localhost:9096,localhost:9097"
	KafkaProducerReturnSuccesses = true
	KafkaConsumerSleep           = 10 * time.Second
	KafkaConsumerReturnErrors    = true
	KafkaTopic                   = "bb_test_topic"
	KafkaGroupId                 = "bb_test_group"
)
