package ports

import (
	"time"
)

type Storage interface {
	SelectByTimeAndCount(start, end time.Time) ([]*Record, error)
}

type Cache interface {
	Set(Key, Value string) error
	Get(key string) (string, error)
}
