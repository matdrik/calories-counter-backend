package service

import (
	calories_counter_backend "server"
	"server/internal/repository"
)

type Authorization interface {
	CreateUser(user calories_counter_backend.User) (int, error)
	GenerateToken(username, password string) (string, error)
	// ParseToken todo - возвращать еще и роль пользователя
	ParseToken(token string) (int, error)
}

type Log interface {
	Create(userId int, log calories_counter_backend.Log) (int, error)
	GetAll(userId int, date string) ([]calories_counter_backend.LogResponse, error)
	GetById(userId, logId int) (calories_counter_backend.LogResponse, error)
	Delete(userId, logId int) error
	Update(userId, logId int, input calories_counter_backend.UpdateLogInput) error
}

type Food interface {
	GetAll() ([]calories_counter_backend.FoodResponse, error)
}

type Service struct {
	Authorization
	Log
	Food
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Log:           NewLogService(repo.Log),
		Food:          NewFoodService(repo.Food),
	}
}
