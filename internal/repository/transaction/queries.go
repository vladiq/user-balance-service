package transaction

import "github.com/jmoiron/sqlx"

type Queries struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Queries {
	return &Queries{DB: db}
}
