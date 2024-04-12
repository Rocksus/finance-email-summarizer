package rest

import (
	"fmt"

	"github.com/Rocksus/fundtract/internal/common/config"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// API defines the methods available for REST Delivery
type API interface {
	RegisterRoutes(router *echo.Group)
}

// Serve will start serving the rest api service
func Serve(cfg *config.Config, srv *echo.Echo) *echo.Echo {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.Start(fmt.Sprint(":", cfg.App.Port)); err != nil {
			log.Fatal().
				Err(err).
				Msg("HTTP server error")
		}
	}()

	log.Info().
		Int("port", cfg.App.Port).
		Msg("HTTP Server is running")

	return srv
}
