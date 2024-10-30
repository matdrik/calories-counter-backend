package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"server/internal/models"
)

func (repo *PGRepo) CreateUser(username, password string) error {
	const queryString = `INSERT INTO users(username, password_hash) VALUES ($1, $2)`
	_, err := repo.pool.Exec(context.Background(), queryString, username, password)
	if err != nil {
		return fmt.Errorf("не удалось добавить пользователя: %w", err)
	}
	return nil
}

func (repo *PGRepo) GetUserByUsername(username string) (models.User, error) {
	const queryString string = "SELECT id, username, password_hash, role_id FROM users WHERE username = $1"

	var data models.User
	err := repo.pool.QueryRow(context.Background(), queryString, username).Scan(&data.ID, &data.Username, &data.PasswordHash, &data.RoleID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return data, nil
}

func (repo *PGRepo) UsernameExists(username string) (bool, error) {
	const queryString string = "SELECT username FROM users WHERE username = $1"

	var data models.User
	err := repo.pool.QueryRow(context.Background(), queryString, username).Scan(&data.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	if data == (models.User{}) {
		return false, nil
	}

	return true, nil
}
