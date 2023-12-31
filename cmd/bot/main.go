package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	repoGRPCPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/grpc_repo"
	repoKafkaPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/kafka_repo"
	repoPostgresPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/postgres"
	repoRedis "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/redis_repo"
	kafkaConsumerPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka/custom_consumer"
	kafkaProducerPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka/custom_sync_producer"
	metricPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/metrics"
	redisWrapperPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/redis_wrapper"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel) // TODO change to INFO
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var pool *pgxpool.Pool
	{
		psqlConn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			configPkg.PosgtreSQLHost, configPkg.PosgtreSQLPort,
			configPkg.PosgtreSQLUser, configPkg.PosgtreSQLPassword,
			configPkg.PosgtreSQLDBname,
		)

		pool_, err := pgxpool.Connect(ctx, psqlConn)
		pool = pool_
		if err != nil {
			log.Panic("can't connect to database", err)
		}
		defer pool.Close()

		if err := pool.Ping(ctx); err != nil {
			log.Panic("ping database error", err)
		}

		config := pool.Config()
		config.MaxConnIdleTime = configPkg.PosgtreSQLMaxConnIdleTime
		config.MaxConnLifetime = configPkg.PosgtreSQLMaxConnLifetime
		config.MinConns = configPkg.PosgtreSQLMinConns
		config.MaxConns = configPkg.PosgtreSQLMaxConns
	}

	// prepare data repository
	repoReal := repoPostgresPkg.New(pool)
	go runRepoGRPCServer(ctx, repoReal, configPkg.RepoGRPCServerAddress)

	// prepare kafka
	brokers := strings.Split(configPkg.KafkaBrokers, ",")
	topic := configPkg.KafkaTopic
	groupId := configPkg.KafkaGroupId

	consumer, err := kafkaConsumerPkg.New(
		brokers,
		repoReal,
		groupId,
	)
	if err != nil {
		log.Panic(err)
	}
	go consumer.Run(ctx, []string{topic}, configPkg.KafkaConsumerSleep)

	producer, err := kafkaProducerPkg.New(brokers)
	if err != nil {
		log.Panic(err)
	}

	// prepare business logic
	client := prepareRepoGRPCClient(configPkg.RepoGRPCServerAddress)
	repoGRPC := repoGRPCPkg.New(client)
	repoKafka := repoKafkaPkg.New(*producer, topic)
	repo := repoRedis.New(repoGRPC, repoKafka, redisWrapperPkg.New(
		configPkg.RedisHost,
		configPkg.RedisDb,
		configPkg.RedisPassword,
		configPkg.RedisExpiration,
	))

	bb := bbPkg.New(repo)

	metricManager := metricPkg.NewMetricManager()
	metricManager.RegisterMany(consumer.GetMetrics())
	metricManager.RegisterMany(repo.GetMetrics())
	go metricManager.Run(configPkg.MetricServerHost)

	go runBot(ctx, bb)
	go runREST(ctx)
	runGRPCServer(ctx, bb)
}
