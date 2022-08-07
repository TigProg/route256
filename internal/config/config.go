package config

import "time"

const (
	TelegramBotApiDebug   = true
	TelegramBotApiTimeout = 60

	GRPCClientTarget  = ":8081"
	GRPCServerAddress = ":8081"

	RESTServerAddress = ":8080"

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
)
