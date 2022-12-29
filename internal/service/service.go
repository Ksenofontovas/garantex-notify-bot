package service

import (
	"github.com/Ksenofontovas/garantex-notify-bot/internal/repository"

	tg "github.com/Ksenofontovas/garantex-notify-bot"
)

type Notify interface {
	CreateTrigger(chatId int64, value float64) error
	GetTriggers(chatId int64) ([]float64, error)
	DeleteTrigger(chatId int64, value float64) error
	GetKeys() ([]int64, error)
}

type Garantex interface {
	GetDepth() (resp tg.DepthResponse, err error)
}

type Service struct {
	Notify
	Garantex
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Garantex: NewGarantexService(repos.Garantex),
		Notify:   NewNotifyService(repos.Notify),
	}
}
