package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
	"payment-service/internal/domain"
)

// CreateTransactions godoc
// @Summary Create new transaction
// @Tags order
// @Description Create new transaction
// @Produce json
// @Param input body domain.PaymentInfo true "Payment info"
// @Success 200 {object} transactionInfo "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /payment/create [post]
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
