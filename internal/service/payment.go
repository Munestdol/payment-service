package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
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

func (s *PaymentService) MakePayment(input domain.PaymentInfo) (string, error) {
	massage := fmt.Sprintf("your order for the amount %v has been accepted", input.TotalPrice)
	answerGrps := true

	err := s.ChangeStatusFD(answerGrps, input.OrderId, input.PaymentType)
	return massage, err
}

func (s *PaymentService) CreateTrasactions(input domain.PaymentInfo) (domain.Transaction, error) {
	var transaction domain.Transaction
	var validate = validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Msg("invalid values of fields")
		return transaction, errors.New("invalid values of fields")
	}

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
	err = s.ChangeStatusFD(answerGrps, input.OrderId, input.PaymentType)
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
