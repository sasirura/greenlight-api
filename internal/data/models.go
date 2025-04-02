package data

import (
	"database/sql"
	"errors"
)

var (
	ErrorRecordNotFound = errors.New("record not found")
	ErrEditConflict     = errors.New("edit conflict")
)

type Model struct {
	Movie MovieModel
	Users UserModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		Movie: MovieModel{DB: db},
		Users: UserModel{DB: db},
	}
}
