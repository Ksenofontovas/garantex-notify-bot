package tgbot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Trigger struct {
	ChatId   int   `json:"chat_id"`
	Triggers []int `json:"triggers"`
}

func (tg *TgBot) GetTriggers(chatId int64) ([]float64, error) {
	return tg.services.GetTriggers(chatId)
}

func (tg *TgBot) CreateTrigger(chatId int64, value float64) error {
	return tg.services.CreateTrigger(chatId, value)
}

func (tg *TgBot) DeleteTrigger(chatId int64, value float64) error {
	return tg.services.DeleteTrigger(chatId, value)
}

func (tg *TgBot) SendNotify() error {

	chatsId, err := tg.services.GetKeys()
	if err != nil {
		return err
	}

	depth, err := tg.services.GetDepth()
	if err != nil {
		return err
	}

	price, err := strconv.ParseFloat(depth.Bids[0].Price, 64)
	count := depth.Bids[0].Amount

	for _, chatId := range chatsId {
		triggers, err := tg.services.GetTriggers(chatId)
		if err != nil {
			return err
		}

		for _, trigger := range triggers {
			if trigger <= price {
				msg := tgbotapi.NewMessage(chatId, "")
				msg.Text = fmt.Sprintf("На бирже идет покупка по цене %v в объеме %v", price, count)
				if _, err := tg.botAPI.Send(msg); err != nil {
					return err
				}
				tg.services.DeleteTrigger(chatId, trigger)
			}
		}
	}

	return err
}
