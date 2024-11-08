package service

import (
	calories_counter_backend "server"
	"server/internal/repository"
)

type LogService struct {
	repo repository.Log
}

func NewLogService(repo repository.Log) *LogService {
	return &LogService{repo: repo}
}

func (s *LogService) Create(userId int, log calories_counter_backend.Log) (int, error) {
	return s.repo.Create(userId, log)
}

func (s *LogService) GetAll(userId int, date string) ([]calories_counter_backend.LogResponse, error) {
	return s.repo.GetAll(userId, date)
}

func (s *LogService) GetById(userId, logId int) (calories_counter_backend.LogResponse, error) {
	return s.repo.GetById(userId, logId)
}

func (s *LogService) Delete(userId, logId int) error {
	return s.repo.Delete(userId, logId)
}

func (s *LogService) Update(userId, logId int, update calories_counter_backend.UpdateLogInput) error {
	if err := update.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, logId, update)
}
