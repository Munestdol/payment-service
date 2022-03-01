package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	config "payment-service/configs"
)

func main() {
	cfg, err := config.Init("./configs")
	if err != nil {
		log.Fatal().Err(err).Msg("wrong config variables")
	}
	fmt.Println(cfg)
}
