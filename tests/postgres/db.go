//go:build integration
// +build integration

package postgres

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/tigprog/bus_booking/tests/config"
)

type TDB struct {
	DB *pgxpool.Pool
	sync.Mutex
}

func New(cfg *config.Config) *TDB {
	psqlConn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
	)

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Panic(err)
	}
	return &TDB{pool, sync.Mutex{}}
}

func (db *TDB) Setup(t *testing.T) {
	t.Helper()
	db.Lock()
	db.truncate(context.Background())
}

func (db *TDB) TearDown(t *testing.T) {
	defer db.Unlock()
	db.truncate(context.Background())
}

func (db *TDB) truncate(ctx context.Context) {
	_, err := db.DB.Query(ctx, "TRUNCATE TABLE public.booking")
	if err != nil {
		log.Panic(err)
	}
}
