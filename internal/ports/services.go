package ports

import "time"

type Services interface {
	SelectDataByDateAndRange(start, end time.Time, min, max int) ([]*FetchRespRecord, error)
	SetIntoMemoryDB(key, value string) error
	GetFromMemoryDB(key string) (string, error)
}
