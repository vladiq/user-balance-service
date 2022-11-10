package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/vladiq/user-balance-service/internal/repository/queries"
)

type repository struct {
	*queries.Queries
	DB *sqlx.DB
	l  zerolog.Logger
}

func NewRepository(l zerolog.Logger, db *sqlx.DB) *repository {
	return &repository{
		DB:      db,
		Queries: queries.New(db),
		l:       l,
	}
}
