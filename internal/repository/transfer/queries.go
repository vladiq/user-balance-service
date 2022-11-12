package transfer

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Queries struct {
	DB     *sqlx.DB
	logger zerolog.Logger
}

func NewQueries(db *sqlx.DB, logger zerolog.Logger) *Queries {
	return &Queries{DB: db, logger: logger}
}
