package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

import (
	_ "github.com/lib/pq"
)

const (
	UserTable  = "users"
	FoodTable  = "food"
	RolesTable = "roles"
	MealsTable = "meals"
	LogsTable  = "logs"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
