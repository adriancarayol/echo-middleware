package middleware_test

import (
	"os"

	middleware "github.com/faabiosr/echo-middleware"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
)

// This example registers the ZeroLog middleware with default configuration.
func ExampleZeroLog() {
	e := echo.New()

	// Middleware
	e.Use(middleware.ZeroLog())
}

// This examples registers the ZeroLog middleware with customer configuration.
func ExampleZeroLogWithConfig() {
	e := echo.New()

	// Custom zerolog logger instance
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Middleware
	logConfig := middleware.ZeroLogConfig{
		Logger: logger,
		FieldMap: map[string]string{
			"uri":    "@uri",
			"host":   "@host",
			"method": "@method",
			"status": "@status",
		},
	}

	e.Use(middleware.ZeroLogWithConfig(logConfig))
}

// This example registers the Logrus middleware with default configuration.
func ExampleLogrus() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logrus())
}

// This examples registers the Logrus middleware with customer configuration.
func ExampleLogrusWithConfig() {
	e := echo.New()

	// Custom logrus logger instance
	logger := logrus.New()

	// Middleware
	logConfig := middleware.LogrusConfig{
		Logger: logger,
		FieldMap: map[string]string{
			"uri":    "@uri",
			"host":   "@host",
			"method": "@method",
			"status": "@status",
		},
	}

	e.Use(middleware.LogrusWithConfig(logConfig))
}
