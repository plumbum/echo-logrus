// Package echologrus provides a middleware for echo v2 that logs request details
// via the logrus logging library
// Original from fknsrs.biz/p/echo-logrus
package echologrus

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
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			latency := time.Since(start)
			request := c.Request()

			entry := l.WithFields(logrus.Fields{
				"prefix": name,
				"remote_ip":   request.RealIP(),
				"method":   request.Method(),
				"uri":  request.URI(),
				"status":   c.Response().Status(),
				"bytes_in": request.ContentLength(),
				"bytes_out": c.Response().Size(),
				"latency":  latency.Nanoseconds()/1000,
				"latency_human":  latency.String(),
			})

			if reqID := request.Header().Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			if err != nil {
				entry.Error(err)
				c.Error(err)
			} else {
				entry.Info("done")
			}

			return nil
		}
	}
}
