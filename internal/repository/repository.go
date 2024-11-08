package repository

import (
	"github.com/jmoiron/sqlx"
	calories_counter_backend "server"
)

type Authorization interface {
	CreateUser(user calories_counter_backend.User) (int, error)
	GetUser(username, password string) (calories_counter_backend.User, error)
}

type Log interface {
	Create(userId int, log calories_counter_backend.Log) (int, error)
	GetAll(userId int, date string) ([]calories_counter_backend.LogResponse, error)
	GetById(userId, logId int) (calories_counter_backend.LogResponse, error)
	Delete(userId, logId int) error
	Update(userId, logId int, update calories_counter_backend.UpdateLogInput) error
}

type Food interface {
	GetAll() ([]calories_counter_backend.FoodResponse, error)
}

type Repository struct {
	Authorization
	Log
	Food
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Log:           NewLogPostgres(db),
		Food:          NewFoodPostgres(db),
	}
}
