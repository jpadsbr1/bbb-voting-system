package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(postgres_url string) *Postgres {
	dsn, err := pgxpool.ParseConfig(postgres_url)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	log.Printf("Connecting to Postgres...")
	return &Postgres{pool: pool}
}

func (p *Postgres) Close() {
	log.Printf("Closing Postgres...")
	p.pool.Close()
}
