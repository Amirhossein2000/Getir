package ports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Record struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Key       string             `bson:"key,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	Counts    []int              `bson:"counts,omitempty"`

	TotalCount int
}

type FetchRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type FetchResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"msg"`
	Records []*FetchRespRecord `json:"records"`
}

type FetchRespRecord struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"created_at"`
	TotalCount int       `json:"total_count"`
}
