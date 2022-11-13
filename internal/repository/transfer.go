package repository

import (
	"github.com/jmoiron/sqlx"
)

type transferRepository struct {
	DB *sqlx.DB
}

func NewTransferRepository(db *sqlx.DB) *transferRepository {
	return &transferRepository{
		DB: db,
	}
}
