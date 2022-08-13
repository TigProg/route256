package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	repoPostgresPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/postgres"
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
			log.Fatal("can't connect to database", err)
		}
		defer pool.Close()

		if err := pool.Ping(ctx); err != nil {
			log.Fatal("ping database error", err)
		}

		config := pool.Config()
		config.MaxConnIdleTime = configPkg.PosgtreSQLMaxConnIdleTime
		config.MaxConnLifetime = configPkg.PosgtreSQLMaxConnLifetime
		config.MinConns = configPkg.PosgtreSQLMinConns
		config.MaxConns = configPkg.PosgtreSQLMaxConns
	}

	repo := repoPostgresPkg.New(pool)
	//repo := repoLocalPkg.New()  // for local cache

	bb := bbPkg.New(repo)

	go runBot(ctx, bb)
	go runREST(ctx)
	runGRPCServer(ctx, bb)
}
