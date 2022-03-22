package handler

import (
	"github.com/gin-gonic/gin"
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
	//var validate = validator.New()
	//if err := validate.Struct(input); err != nil {
	//	log.Error().Err(err).Msg("invalid values of fields")
	//	newResponse(c, http.StatusBadRequest, "invalid values of fields")
	//	return
	//}
	transactInfo, err := h.services.Payment.CreateTrasactions(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if transactInfo.Status == "canceled" {
		newResponse(c, http.StatusBadRequest, "transaction - canceled")
		return
	}
	c.JSON(http.StatusOK, transactionInfo{
		Data: transactInfo,
	})
}

func (h *Handler) MakePayment(c *gin.Context) {
	var input domain.PaymentInfo
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	switch input.PaymentType {
	case "cardonline":
		transactInfo, err := h.services.Payment.CreateTrasactions(input)
		if err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if transactInfo.Status == "canceled" {
			newResponse(c, http.StatusBadRequest, "transaction - canceled")
			return
		}
		c.JSON(http.StatusOK, transactionInfo{
			Data: transactInfo,
		})
	case "card":
		transactInfo, err := h.services.Payment.MakePayment(input)
		if err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, payInfo{
			Data: transactInfo,
		})
	case "cash":
		transactInfo, err := h.services.Payment.MakePayment(input)
		if err != nil {
			newResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, payInfo{
			Data: transactInfo,
		})
	}
}
