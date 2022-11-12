package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

// Store provides all functions to execute db queries and transactions
//type Store struct {
//	accountQueries     *account.Queries
//	reservationQueries *reservation.Queries
//	transactionQueries *transaction.Queries
//
//	db     *sqlx.DB
//	logger zerolog.Logger
//}
//
//func NewStore(db *sqlx.DB, logger zerolog.Logger) *Store {
//	return &Store{
//		accountQueries:     account.NewQueries(db, logger),
//		reservationQueries: reservation.NewQueries(db, logger),
//		transactionQueries: transaction.NewQueries(db, logger),
//		db:                 db,
//		logger:             logger,
//	}
//}
//
//func (store *Store) execTx(ctx context.Context)

type MakeMoneyTransactionParams struct {
	fromID
	toID
}

func MakeMoneyTransaction(ctx context.Context, db *sqlx.DB, logger zerolog.Logger) {

}
