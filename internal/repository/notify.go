package repository

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/recoilme/pudge"
)

type NotifyPudge struct {
	db *pudge.Db
}

func NewNotifyPudge(db *pudge.Db) *NotifyPudge {
	return &NotifyPudge{db: db}
}

func (r *NotifyPudge) CreateTrigger(chatId int64, value float64) error {
	var values []float64

	r.db.Get(chatId, &values)
	for i := 0; i < len(values); i++ {
		if value == values[i] {
			strVal := fmt.Sprintf("%.2f", value)
			return errors.New("Триггер со значением " + strVal + " уже установлен!")
		}
	}

	values = append(values, value)
	return r.db.Set(chatId, values)
}

func (r *NotifyPudge) GetTriggers(chatId int64) ([]float64, error) {
	var values []float64
	err := r.db.Get(chatId, &values)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (r *NotifyPudge) DeleteTrigger(chatId int64, value float64) error {
	var values []float64

	if err := r.db.Get(chatId, &values); err != nil {
		return err
	}
	for i := 0; i < len(values); i++ {
		if value == values[i] {
			values = append(values[:i], values[i+1:]...)
			if err := r.db.Set(chatId, values); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("триггер не найден")
}

func (r *NotifyPudge) GetKeys() ([]int64, error) {
	var values []int64

	keys, err := r.db.Keys(0, 0, 0, true)

	for i := 0; i < len(keys); i++ {
		val := binary.BigEndian.Uint64(keys[i])
		values = append(values, int64(val))
	}

	return values, err
}
