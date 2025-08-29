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
	pool, err := pgxpool.New(context.Background(), postgres_url)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connecting to Postgres...")
	return &Postgres{pool: pool}
}

func (p *Postgres) Close() {
	log.Printf("Closing Postgres...")
	p.pool.Close()
}
