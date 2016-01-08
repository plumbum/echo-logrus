// Package echologrus provides a middleware for echo that logs request details
// via the logrus logging library
package echologrus // fknsrs.biz/p/echo-logrus

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// New returns a new middleware handler with a default name and logger
func New() echo.MiddlewareFunc {
	return NewWithName("web")
}

// NewWithName returns a new middleware handler with the specified name
func NewWithName(name string) echo.MiddlewareFunc {
	return NewWithNameAndLogger(name, logrus.StandardLogger())
}

// NewWithNameAndLogger returns a new middleware handler with the specified name
// and logger
func NewWithNameAndLogger(name string, l *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()
			isError := false

			if err := next(c); err != nil {
				c.Error(err)
				isError = true
			}

			latency := time.Since(start)

			entry := l.WithFields(logrus.Fields{
				"request": c.Request().RequestURI,
				"method":  c.Request().Method,
				"remote":  c.Request().RemoteAddr,
				"status":  c.Response().Status(),
				"latency": latency,
			})

			if reqID := c.Request().Header.Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			if isError {
				entry.Error("error by handling request")
			} else {
				entry.Info("completed handling request")
			}

			return nil
		}
	}
}
