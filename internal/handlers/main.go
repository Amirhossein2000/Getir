package handlers

import (
	"Getir/internal/ports"
)

type handlers struct {
	services ports.Services
}

func NewHandlers(services ports.Services) *handlers {
	return &handlers{
		services: services,
	}
}
