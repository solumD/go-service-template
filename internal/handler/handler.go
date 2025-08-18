package handler

import (
	"github.com/solumD/go-service-template/internal/service"
)

type handler struct {
	service service.EntityService
}

func New(s service.EntityService) *handler {
	return &handler{
		service: s,
	}
}
