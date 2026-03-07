package dto

import "github.com/solumD/go-service-template/internal/model"

func FromCreateEntityReqToModel(req CreateEntityReq) *model.Entity {
	return &model.Entity{
		Name: req.Name,
	}
}

func FromEntityModelToResp(entity *model.Entity) *GetEntityByIDResp {
	return &GetEntityByIDResp{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
