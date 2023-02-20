package handlers

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

type InMemoryData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorResp struct {
	Code    int
	Message string
}

func (h *handlers) InMemory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    4,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		})
	}
}

func (h *handlers) handlePost(w http.ResponseWriter, r *http.Request) {
	req := InMemoryData{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    1,
			Message: "Invalid request payload",
		})
		return
	}

	if req.Key == "" || req.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    1,
			Message: "Key and Value are required",
		})
		return
	}

	err = h.services.SetIntoMemoryDB(req.Key, req.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    2,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		log.Printf("setIntoMemoryDB has failed. err: %s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func (h *handlers) handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    1,
			Message: "Missing key parameter",
		})
		return
	}

	val, err := h.services.GetFromMemoryDB(key)
	if err != nil {
		if err == redis.Nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResp{
				Code:    3,
				Message: http.StatusText(http.StatusNotFound),
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResp{
				Code:    2,
				Message: http.StatusText(http.StatusInternalServerError),
			})
			log.Printf("getFromMemoryDB has failed. err: %s\n", err.Error())
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(InMemoryData{
		Key:   key,
		Value: val,
	})
}
