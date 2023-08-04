package handler

import "github.com/mereiamangeldin/One-lab-Homework-1/internal/service"

type Handler struct {
	srvs service.Service
}

func New(srvs service.Service) *Handler {
	return &Handler{
		srvs: srvs,
	}
}
