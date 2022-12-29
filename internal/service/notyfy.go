package service

import "github.com/Ksenofontovas/garantex-notify-bot/internal/repository"

type NotifyService struct {
	repo repository.Notify
}

func NewNotifyService(repo repository.Notify) *NotifyService {
	return &NotifyService{repo: repo}
}

func (s *NotifyService) CreateTrigger(chatId int64, value float64) error {
	return s.repo.CreateTrigger(chatId, value)
}

func (s *NotifyService) GetTriggers(chatId int64) ([]float64, error) {
	return s.repo.GetTriggers(chatId)
}

func (s *NotifyService) DeleteTrigger(chatId int64, value float64) error {
	return s.repo.DeleteTrigger(chatId, value)
}

func (s *NotifyService) GetKeys() ([]int64, error) {
	return s.repo.GetKeys()
}
