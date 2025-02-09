package router

import (
	"github.com/Rocksus/fundtract/internal/delivery/http/handler"
	"github.com/Rocksus/fundtract/internal/platform/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes(e *echo.Echo) {
	e.GET("/health", handler.HealthCheck)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		// LogValuesFunc is called after the request is completed.
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.InfoWithFields(
				"incoming request",
				log.KV{
					"method":     v.Method,
					"url":        v.URI,
					"status":     v.Status,
					"latency":    v.Latency,
					"remote_ip":  v.RemoteIP,
					"user_agent": v.UserAgent,
				},
			)
			return nil
		},
	}))
}
