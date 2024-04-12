package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/Rocksus/fundtract/internal/common/config"
	"github.com/Rocksus/fundtract/internal/delivery/rest"
	"github.com/Rocksus/fundtract/internal/delivery/rest/handler"
)

//go:generate sqlc generate -f ../../sqlc.yaml

func main() {
	env := config.GetEnv()

	cfg, err := config.Init("files/etc/fundtract/", env)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed initializing config")
	}

	e := echo.New()

	handler.RegisterHandlers(e, []rest.API{})

	if err := e.Start(fmt.Sprint(":", cfg.App.Port)); err != nil {
		log.Fatal().Err(err).Msg("HTTP Server Error")
	}
}
