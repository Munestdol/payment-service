package main

import (
	"log"
	config "payment-service/configs"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Init("./configs")
	if err != nil {
		log.Fatal().Err(err).Msg("wrong config variables")
	}
}
