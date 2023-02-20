package handlers

import (
	"Getir/internal/services"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInMemory(t *testing.T) {
	data := make(map[string]string)
	cache := &mockCache{data: data}
	s := services.NewServices(nil, cache)
	h := NewHandlers(s)

	testInMemoryData := &InMemoryData{
		Key:   "testKey",
		Value: "testValue",
	}
	testTablePost := []struct {
		request      *InMemoryData
		expectedResp *InMemoryData
	}{
		{
			request:      testInMemoryData,
			expectedResp: testInMemoryData,
		},
	}

	for _, tt := range testTablePost {
		resp := InMemoryData{}
		err := doTestRequest(http.MethodPost, "/in-memory", h.InMemory, tt.request, &resp)
		if err != nil {
			t.Errorf("in-memory request test faild: %s", err.Error())
		}
		assert.Equal(t, tt.expectedResp.Key, resp.Key, "key is not equal")
		assert.Equal(t, tt.expectedResp.Value, resp.Value, "value is not equal")
	}

	testTableGet := []struct {
		request         string
		expectedResp    *InMemoryData
		expectedErrResp *ErrorResp
	}{
		{
			expectedErrResp: &ErrorResp{
				Code:    1,
				Message: "Missing key parameter",
			},
		},
		{
			request: "doesNotExist",
			expectedErrResp: &ErrorResp{
				Code:    3,
				Message: http.StatusText(http.StatusNotFound),
			},
		},
		{
			request: testTablePost[0].request.Key,
			expectedResp: &InMemoryData{
				Key:   testTablePost[0].request.Key,
				Value: testTablePost[0].request.Value,
			},
		},
	}

	for _, tt := range testTableGet {
		if tt.expectedErrResp != nil {
			resp := ErrorResp{}
			err := doTestRequest(http.MethodGet, fmt.Sprintf("/in-memory?key=%s", tt.request), h.InMemory, nil, &resp)
			if err != nil {
				t.Errorf("in-memory request test faild: %s", err.Error())
			}
			assert.Equal(t, tt.expectedErrResp.Code, resp.Code, "error code is not equal")
			assert.Equal(t, tt.expectedErrResp.Message, resp.Message, "error message is not equal")
			continue
		}

		resp := InMemoryData{}
		err := doTestRequest(http.MethodGet, fmt.Sprintf("/in-memory?key=%s", tt.request), h.InMemory, nil, &resp)
		if err != nil {
			t.Errorf("in-memory request test faild: %s", err.Error())
		}
		assert.Equal(t, tt.expectedResp.Key, resp.Key, "key is not equal")
		assert.Equal(t, tt.expectedResp.Value, resp.Value, "value is not equal")
	}
}

type mockCache struct {
	data map[string]string
}

func (m *mockCache) Set(Key, Value string) error {
	m.data[Key] = Value
	return nil
}

func (m *mockCache) Get(key string) (string, error) {
	value, ok := m.data[key]
	if !ok {
		return "", redis.Nil
	}
	return value, nil
}
