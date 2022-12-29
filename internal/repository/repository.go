package repository

import (
	"tg"
	"time"

	"github.com/recoilme/pudge"
)

type Garantex interface {
	GetDepth() (tg.DepthResponse, error)
}

type Notify interface {
	CreateTrigger(chatId int64, value float64) error
	GetTriggers(chatId int64) ([]float64, error)
	DeleteTrigger(chatId int64, value float64) error
	GetKeys() ([]int64, error)
}

type Repository struct {
	Garantex
	Notify
}

func NewRepository(host string, timeout time.Duration, db *pudge.Db) *Repository {
	return &Repository{
		Garantex: NewClient(host, timeout),
		Notify:   NewNotifyPudge(db),
	}
}
