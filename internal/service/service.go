package service

import (
	config "payment-service/configs"
	"payment-service/internal/domain"
	"payment-service/internal/repository"
)

type Payment interface {
	CreateTrasactions(input domain.PaymentInfo) (domain.Transaction, error)
	BoolStatus(input domain.Transaction) bool
}

type Service struct {
	Payment
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Payment: NewPaymentService(repos.Payment, cfg),
	}
}
