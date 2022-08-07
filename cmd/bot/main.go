package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	repoPostgresPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/postgres"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "user"
	Password = "password"
	DBname   = "bus_booking"

	MaxConnIdleTime = time.Minute
	MaxConnLifetime = time.Hour
	MinConns        = 2
	MaxConns        = 4
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var pool *pgxpool.Pool
	{
		psqlConn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			Host, Port, User, Password, DBname,
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
		config.MaxConnIdleTime = MaxConnIdleTime
		config.MaxConnLifetime = MaxConnLifetime
		config.MinConns = MinConns
		config.MaxConns = MaxConns
	}

	repo := repoPostgresPkg.New(pool)
	//repo := repoLocalPkg.New()  // for local cache

	bb := bbPkg.New(repo)

	go runBot(ctx, bb)
	go runREST(ctx)
	runGRPCServer(ctx, bb)
}
