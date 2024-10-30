package migrate

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
)

func RunMigrations() error {
	//conf := config.New()
	//postgresql://postgres:8563@localhost:5432/calories_counter
	db, err := sql.Open("postgres", "user=postgres password=8563 dbname=calories_counter sslmode=disable")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("could not start migration: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not apply migration: %w", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
