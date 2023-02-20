package help

import (
	"Getir/internal/ports"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const TimeFormat = "2006-01-02"

func FilterRecords(records []*ports.Record, min, max int) []*ports.FetchRespRecord {
	respRecords := make([]*ports.FetchRespRecord, 0, len(records))
	for _, r := range records {
		totalCount := 0
		for _, c := range r.Counts {
			totalCount += c
		}
		if totalCount > min && totalCount < max {
			respRecords = append(respRecords, &ports.FetchRespRecord{
				Key:        r.Key,
				CreatedAt:  r.CreatedAt,
				TotalCount: totalCount,
			})
		}
	}

	return respRecords
}

func TestTime(strTime string) time.Time {
	t, _ := time.Parse(TimeFormat, strTime)
	return t
}

func AssertRecords(t *testing.T, gotRecords, expectedRecords []*ports.FetchRespRecord) {
	assert.Equal(t, len(gotRecords), len(expectedRecords), "length of records is not equal")
	for i := range gotRecords {
		assert.Equal(t, gotRecords[i].Key, expectedRecords[i].Key, "Key record is not equal")
		assert.Equal(t, gotRecords[i].CreatedAt, expectedRecords[i].CreatedAt, "CreatedAt record are not equal")
		assert.Equal(t, gotRecords[i].TotalCount, expectedRecords[i].TotalCount, "TotalCount record is not equal")
	}
}
