package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/solumD/go-service-template/internal/delivery/http/v1/dto"
	"github.com/solumD/go-service-template/pkg/logger"
)

var (
	ErrFailedToCreateEntity  = errors.New("failed to create entity")
	ErrInvalidEntityIDType   = errors.New("invalid entity id type")
	ErrFailedToGetEntityByID = errors.New("failed to get entity by id")
)

func (h *handler) createEntity(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handler.CreateEntity"
		log := h.log.With(logger.String("fn", fn))

		log.Info("new request")

		var req dto.CreateEntityReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", logger.Error(err))

			h.errorResponse(w, contentTypeJSON, http.StatusBadRequest, ErrFailedToDecodeReq)
			return
		}

		id, err := h.entityUsecase.CreateEntity(ctx, dto.FromCreateEntityReqToModel(req))
		if err != nil {
			log.Error("failed to create entity", logger.Error(err))

			h.errorResponse(w, contentTypeJSON, http.StatusInternalServerError, ErrFailedToCreateEntity)
			return
		}

		log.Info("created entity", logger.Int("entity id", id))

		h.response(w, contentTypeEmpty, http.StatusCreated, nil)
	}
}

func (h *handler) getEntityByID(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handler.GetEntityByID"
		log := h.log.With(logger.String("fn", fn))

		log.Info("new request")

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to get entity id", logger.Error(err))

			h.errorResponse(w, contentTypeJSON, http.StatusBadRequest, ErrInvalidEntityIDType)
			return
		}

		entity, err := h.entityUsecase.GetEntityByID(ctx, id)
		if err != nil {
			log.Error("failed to get entity by id", logger.Error(err))

			h.errorResponse(w, contentTypeJSON, http.StatusInternalServerError, ErrFailedToGetEntityByID)
			return
		}

		resp := dto.FromEntityModelToResp(entity)
		respBody, err := json.Marshal(resp)
		if err != nil {
			log.Error("failed to marshal response", logger.Error(err))

			h.errorResponse(w, contentTypeJSON, http.StatusInternalServerError, ErrFailedToGetEntityByID)
			return
		}

		log.Info("got entity by id", logger.Any("entity", entity))

		h.response(w, contentTypeJSON, http.StatusOK, respBody)
	}
}
