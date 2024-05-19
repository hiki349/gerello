package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PG struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, connStr string, maxAttempts int) (*PG, error) {
	var pool *pgxpool.Pool
	var err error
log.Println(maxAttempts)
	for maxAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		pool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			log.Println("1")
			maxAttempts--
			continue
		}

		break
	}
	log.Println(maxAttempts)

	return &PG{
		Pool: pool,
	}, err
}