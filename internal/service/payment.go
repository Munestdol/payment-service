package service

import (
	"github.com/rs/zerolog/log"
	config "payment-service/configs"
	"payment-service/internal/domain"
	"payment-service/internal/repository"
	"strconv"
)

type PaymentService struct {
	repo repository.Payment
	cfg  *config.Config
}

func NewPaymentService(repo repository.Payment, cfg *config.Config) *PaymentService {
	return &PaymentService{repo: repo, cfg: cfg}
}

func (s *PaymentService) CreateTrasactions(input domain.PaymentInfo) (domain.Transaction, error) {
	status := s.GetPaymentStatus(input.CVV)
	transaction, err := s.repo.CreateTrasactions(status, input)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	cardNumber := s.ChangeCardNumber(input.CardNumber)
	transaction.CardNumber = cardNumber
	transaction.Status = status
	transaction.TotalPrice = input.TotalPrice
	answerGrps := s.BoolStatus(transaction)
	err = s.ChangeStatusFD(answerGrps, input.OrderId)

	return transaction, err
}

func (s *PaymentService) GetPaymentStatus(cvv string) string {
	ccv, _ := strconv.Atoi(cvv)
	var status string
	if ccv%2 == 0 {
		status = "canceled"
	} else {
		status = "approved"
	}
	return status
}

func (s *PaymentService) ChangeCardNumber(number string) string {
	newNumber := "**** **** **** " + number[12:]
	return newNumber
}

func (s *PaymentService) BoolStatus(input domain.Transaction) bool {
	if input.Status == "canceled" {
		return false
	}
	return true
}
