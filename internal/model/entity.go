package model

type Entity struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}
