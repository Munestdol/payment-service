package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"payment-service/internal/domain"
)

type transactionInfo struct {
	Data domain.Transaction `json:"data"`
}

type payInfo struct {
	Data string `json:"data"`
}
type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	log.Error().Msg(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
