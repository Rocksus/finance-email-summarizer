package handler

import (
	"github.com/Rocksus/fundtract/internal/delivery/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func RegisterHandlers(e *echo.Echo, handlers []rest.API) {
	log.Info().Msg("Registering handlers")

	e.GET("/", HealthCheck)

	apiRoute := e.Group("/api")

	// apply CORS to all api endpoints
	apiRoute.Use(
		middleware.CORS(),
	)

	for _, h := range handlers {
		h.RegisterRoutes(apiRoute)
	}
}
