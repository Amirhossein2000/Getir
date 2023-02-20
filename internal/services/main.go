package services

import (
	"Getir/internal/help"
	"Getir/internal/ports"
	"time"
)

type services struct {
	storage ports.Storage
	cache   ports.Cache
}

func NewServices(storage ports.Storage, cache ports.Cache) ports.Services {
	return &services{
		storage: storage,
		cache:   cache,
	}
}

func (s *services) SelectDataByDateAndRange(start, end time.Time, min, max int) ([]*ports.FetchRespRecord, error) {
	records, err := s.storage.SelectByTimeAndCount(start, end)
	if err != nil {
		return nil, err
	}

	return help.FilterRecords(records, min, max), nil
}

func (s *services) SetIntoMemoryDB(key, value string) error {
	return s.cache.Set(key, value)
}

func (s *services) GetFromMemoryDB(key string) (string, error) {
	return s.cache.Get(key)
}
