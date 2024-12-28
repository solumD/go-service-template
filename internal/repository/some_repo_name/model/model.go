package model

// for repository DTOS

// SomeModel model of some object from repo
type SomeModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
