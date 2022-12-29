package repository

import (
	"github.com/recoilme/pudge"
)

func NewPudgeDB(location string, syncInterval int) (*pudge.Db, error) {
	conf := &pudge.Config{
		SyncInterval: syncInterval} // every second fsync
	db, err := pudge.Open(location, conf) //("./db", cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}
