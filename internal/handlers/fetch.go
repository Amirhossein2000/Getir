package handlers

import (
	"Getir/internal/help"
	"Getir/internal/ports"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (h *handlers) Fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResp{
			Code:    4,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	req := ports.FetchRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    1,
			Message: "Invalid request payload",
		})
		return
	}

	if req.MinCount < 0 || req.MaxCount < 0 || req.MinCount > req.MaxCount {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    1,
			Message: "Invalid minCount and maxCount values",
		})
		return
	}

	if req.StartDate == "" || req.EndDate == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    1,
			Message: "Start date and end date are required",
		})
		return
	}

	startDate, err := time.Parse(help.TimeFormat, req.StartDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    1,
			Message: "Invalid startDate",
		})
		return
	}
	endDate, err := time.Parse(help.TimeFormat, req.EndDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    1,
			Message: "Invalid endDate",
		})
		return
	}

	records, err := h.services.SelectDataByDateAndRange(startDate, endDate, req.MinCount, req.MaxCount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    2,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		log.Printf("fetch has failed. err: %s\n", err.Error())
		return
	}

	if len(records) < 1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ports.FetchResponse{
			Code:    3,
			Message: http.StatusText(http.StatusNotFound),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ports.FetchResponse{
		Message: "success",
		Records: records,
	})
}
