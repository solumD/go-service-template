package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/pkg/logger"
	"go.uber.org/zap"
)

type CreateEntityReq struct {
	model.Entity
}

type CreateEntityResp struct {
	ID int64 `json:"id"`
}

func (h *handler) CreateEntity(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateEntityReq
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			logger.Error("[handler] failed to unmarshal request",
				zap.Error(err),
			)

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, NewErrorResponse("failed to unmarshal request"))
			return
		}

		id, err := h.service.CreateEntity(ctx, &req.Entity)
		if err != nil {
			logger.Error("[handler] failed to create entity",
				zap.Error(err),
			)

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, NewErrorResponse("failed to create entity"))
			return
		}

		logger.Info("[handler] created entity",
			zap.Int64("id", id),
		)

		render.Status(r, http.StatusOK)
		render.JSON(w, r, CreateEntityResp{
			ID: id,
		})
	}
}

type GetEntityResp struct {
	model.Entity
}

func (h *handler) GetEntity(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r.Context())

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logger.Error("[handler] failed to parse id",
				zap.Error(err),
				zap.String("request_id",
					reqID,
				))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, NewErrorResponse("failed to parse id"))
			return
		}

		entity, err := h.service.GetEntity(ctx, int64(id))
		if err != nil {
			logger.Error("[handler] failed to get entity",
				zap.Error(err),
				zap.String("request_id",
					reqID,
				))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, NewErrorResponse("failed to get entity"))
			return
		}

		logger.Info("[handler] got entity",
			zap.Int64("id", entity.ID),
			zap.String("request_id", reqID),
		)

		render.Status(r, http.StatusOK)
		render.JSON(w, r, GetEntityResp{
			Entity: *entity,
		})
	}
}
