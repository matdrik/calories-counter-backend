package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	calories_counter_backend "server"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user calories_counter_backend.User) (int, error) {
	//return 1111, nil

	var id int
	queryString := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", UserTable)
	row := r.db.QueryRow(context.Background(), queryString, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (calories_counter_backend.User, error) {
	var user calories_counter_backend.User
	queryString := fmt.Sprintf("select id, username, password_hash from %s where username=$1 and password_hash=$2", UserTable)
	err := r.db.QueryRow(context.Background(), queryString, username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
