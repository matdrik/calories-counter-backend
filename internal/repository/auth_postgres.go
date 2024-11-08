package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	calories_counter_backend "server"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user calories_counter_backend.User) (int, error) {
	var id int
	queryString := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", UserTable)
	row := r.db.QueryRow(queryString, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (calories_counter_backend.User, error) {
	var user calories_counter_backend.User
	queryString := fmt.Sprintf("select id from %s where username=$1 and password_hash=$2", UserTable)
	err := r.db.Get(&user, queryString, username, password)

	return user, err
}
