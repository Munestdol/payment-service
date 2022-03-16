package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"payment-service/internal/domain"
	"time"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) CreateTrasactions(status string, input domain.PaymentInfo) (domain.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	var transaction domain.Transaction
	date := time.Now()
	query := fmt.Sprintf(`INSERT INTO transactions (user_id, order_id, date, cost,status)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`)
	var id string
	row := r.db.QueryRow(query, input.UserId, input.OrderId, date, input.TotalPrice, status)
	if err = row.Scan(&id); err != nil {

		tx.Rollback()
		log.Error().Err(err).Msg("")
	}
	transaction.Id = id
	transaction.Date = date.Format("02-01-2006, 15:04")
	return transaction, tx.Commit()
}
