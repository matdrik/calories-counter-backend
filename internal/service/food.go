package service

import (
	calories_counter_backend "server"
	"server/internal/repository"
)

type FoodService struct {
	repo repository.Food
}

func NewFoodService(repo repository.Food) *FoodService {
	return &FoodService{repo: repo}
}

func (s *FoodService) GetAll() ([]calories_counter_backend.FoodResponse, error) {
	return s.repo.GetAll()
}
