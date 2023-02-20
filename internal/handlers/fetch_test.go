package handlers

import (
	"Getir/internal/help"
	"Getir/internal/ports"
	"Getir/internal/services"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	data := []*ports.Record{
		{
			Key:        "test1",
			CreatedAt:  help.TestTime("2016-01-28"),
			Counts:     []int{10, 20, 30},
			TotalCount: 60,
		},
	}
	storage := &mockStorage{data: data}
	s := services.NewServices(storage, nil)
	h := NewHandlers(s)

	testTable := []struct {
		request      *ports.FetchRequest
		expectedResp *ports.FetchResponse
	}{
		{
			request: &ports.FetchRequest{
				StartDate: "2016-01-26",
				EndDate:   "2018-02-02",
				MinCount:  -1,
				MaxCount:  -100,
			},
			expectedResp: &ports.FetchResponse{
				Code:    1,
				Message: "Invalid minCount and maxCount values",
			},
		},
		{
			request: &ports.FetchRequest{
				MinCount: 10,
				MaxCount: 20,
			},
			expectedResp: &ports.FetchResponse{
				Code:    1,
				Message: "Start date and end date are required",
			},
		},
		{
			request: &ports.FetchRequest{
				StartDate: "2016/01/26",
				EndDate:   "2018-02-02",
			},
			expectedResp: &ports.FetchResponse{
				Code:    1,
				Message: "Invalid startDate",
			},
		},
		{
			request: &ports.FetchRequest{
				StartDate: "2016-01-26",
				EndDate:   "2018/02/02",
			},
			expectedResp: &ports.FetchResponse{
				Code:    1,
				Message: "Invalid endDate",
			},
		},
		{
			request: &ports.FetchRequest{
				StartDate: "2003-01-26",
				EndDate:   "2004-02-02",
				MinCount:  1,
				MaxCount:  2900,
			},
			expectedResp: &ports.FetchResponse{
				Code:    3,
				Message: http.StatusText(http.StatusNotFound),
			},
		},
		{
			request: &ports.FetchRequest{
				StartDate: "2016-01-26",
				EndDate:   "2018-02-02",
				MinCount:  1,
				MaxCount:  2900,
			},
			expectedResp: &ports.FetchResponse{
				Code:    0,
				Message: "success",
				Records: []*ports.FetchRespRecord{
					{
						Key:        data[0].Key,
						CreatedAt:  data[0].CreatedAt,
						TotalCount: data[0].TotalCount,
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		resp := ports.FetchResponse{}
		err := doTestRequest(http.MethodPost, "/fetch", h.Fetch, tt.request, &resp)
		if err != nil {
			t.Errorf("fetch request test faild: %s", err.Error())
		}
		assert.Equal(t, tt.expectedResp.Code, resp.Code, "code is not equal")
		assert.Equal(t, tt.expectedResp.Message, resp.Message, "message is not equal")
		help.AssertRecords(t, resp.Records, tt.expectedResp.Records)
	}
}

func doTestRequest(method, url string, h http.HandlerFunc, request interface{}, response interface{}) error {
	body := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(body).Encode(request)
	if err != nil {
		return err
	}
	req := httptest.NewRequest(method, url, body)
	recorder := httptest.NewRecorder()
	h(recorder, req)
	return json.NewDecoder(recorder.Result().Body).Decode(response)
}

type mockStorage struct {
	data []*ports.Record
}

func (m *mockStorage) SelectByTimeAndCount(start, end time.Time) ([]*ports.Record, error) {
	output := []*ports.Record{}
	for _, record := range m.data {
		if record.CreatedAt.After(start) && record.CreatedAt.Before(end) {
			output = append(output, record)
		}
	}

	return output, nil
}
