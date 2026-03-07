package dto

type CreateEntityReq struct {
	Name string `json:"name"`
}

type GetEntityByIDResp struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}
