package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	repoGRPCPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/grpc_repo"
	repoPostgresPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/postgres"
	repoRWPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/rw_repo"
	kafkaConsumerPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka/custom_consumer"
	kafkaProducerPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/kafka/custom_sync_producer"
)

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
	go consumer.Run(context.Background(), []string{topic}, configPkg.KafkaConsumerSleep)

	producer, err := kafkaProducerPkg.New(brokers)
	if err != nil {
		log.Panic(err)
	}

	// prepare business logic
	client := prepareRepoGRPCClient(configPkg.RepoGRPCServerAddress)
	repoGRPC := repoGRPCPkg.New(client)
	repo := repoRWPkg.New(repoGRPC, *producer, topic)

	bb := bbPkg.New(repo)

	go runBot(ctx, bb)
	go runREST(ctx)
	runGRPCServer(ctx, bb)
}
