package service

import (
	tg "github.com/Ksenofontovas/garantex-notify-bot"

	"github.com/Ksenofontovas/garantex-notify-bot/internal/repository"
)

type GarantexService struct {
	repo repository.Garantex
}

func NewGarantexService(repo repository.Garantex) *GarantexService {
	return &GarantexService{repo: repo}
}

func (s *GarantexService) GetDepth() (resp tg.DepthResponse, err error) {
	return s.repo.GetDepth()
}
