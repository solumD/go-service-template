package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/solumD/go-service-template/internal/transport/handler/dto"
)

const (
	contentTypeEmpty = ""
	contentTypeJSON  = "application/json"
)

var (
	ErrFailedToDecodeReq = errors.New("failed to decode request")
)

type handler struct {
	entityUsecase EntityUsecase
	log           *slog.Logger
}

func New(uc EntityUsecase, l *slog.Logger) *handler {
	return &handler{
		entityUsecase: uc,
		log:           l,
	}
}

func (h *handler) response(w http.ResponseWriter, contentType string, statusCode int, body []byte) {
	if len(contentType) > 0 {
		w.Header().Add("Content-Type", contentType)
	}

	if statusCode > 0 {
		w.WriteHeader(statusCode)
	}

	if body != nil {
		w.Write(body)
	}
}

func (h *handler) errorResponse(w http.ResponseWriter, contentType string, statusCode int, err error) {
	body, errMarsh := json.Marshal(dto.NewErrorResponse(err.Error()))
	if errMarsh != nil {
		h.response(w, contentTypeEmpty, http.StatusInternalServerError, nil)
		return
	}

	if len(contentType) > 0 {
		w.Header().Add("Content-Type", contentType)
	}

	if statusCode > 0 {
		w.WriteHeader(statusCode)
	}

	w.Write(body)
}
