package repository

import (
	"github.com/jmoiron/sqlx"
	"payment-service/internal/domain"
)

type Payment interface {
	CreateTrasactions(status string, input domain.PaymentInfo) (domain.Transaction, error)
}

type Repository struct {
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Payment: NewPaymentPostgres(db),
	}
}
