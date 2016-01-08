// Package echologrus provides a middleware for echo that logs request details
// via the logrus logging library
package echologrus // fknsrs.biz/p/echo-logrus

import (
	"net"
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
			var err error
			start := time.Now()

			err = next(c)

			latency := time.Since(start)
			request := c.Request()

			remoteIP, _, errSplit := net.SplitHostPort(request.RemoteAddr)
			if errSplit != nil {
				logrus.WithFields(logrus.Fields{
					"func":  "net.SplitHostPort",
					"error": errSplit,
				}).Error("Can't extract remote IP")
				remoteIP = request.RemoteAddr
			}

			entry := l.WithFields(logrus.Fields{
				"request": request.RequestURI,
				"method":  request.Method,
				"protocol": request.Proto,
				"remote":  remoteIP,
				"status":  c.Response().Status(),
				"latency": latency,
			})

			if reqID := request.Header.Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			if err != nil {
				entry.Error(err)
				c.Error(err)
			} else {
				entry.Info("completed handling request")
			}

			return nil
		}
	}
}
