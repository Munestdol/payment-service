package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
	"payment-service/internal/domain"
)

func (h *Handler) CreateTransactions(c *gin.Context) {
	var input domain.PaymentInfo
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	var validate = validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Msg("invalid values of fields")
		newResponse(c, http.StatusBadRequest, "invalid values of fields")
		return
	}
	transactInfo, err := h.services.Payment.CreateTrasactions(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transactionInfo{
		Data: transactInfo,
	})
}
