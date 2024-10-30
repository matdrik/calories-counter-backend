package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type PGRepo struct {
	mu   sync.Mutex
	pool *pgxpool.Pool
}

func New(connectionString string) (*PGRepo, error) {
	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return &PGRepo{mu: sync.Mutex{}, pool: dbpool}, nil
}
