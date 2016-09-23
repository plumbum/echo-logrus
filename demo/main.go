package main

import (
	glog "github.com/labstack/gommon/log"
	log "github.com/plumbum/echo-logrus"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"os"
	"time"
)

type Data struct {
	Id   int
	Name string
	Tags []string
}

func main() {

	logr := logrus.New()
	logr.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: time.StampMilli,
	}

	elog := log.New(logr)
	elog.SetOutput(os.Stderr)
	elog.SetLevel(glog.DEBUG)

	data := glog.JSON{"a": 1, "b": 2}
	elog.Printj(data)
	elog.Debugj(data)
	elog.Infoj(data)
	elog.SetJsonMsg("JSON")
	elog.Warnj(data)
	elog.Errorj(data)

	e := echo.New()
	e.SetDebug(true)
	e.SetLogger(elog)
	e.Use(log.NewMiddlewareWithNameAndLogger("web", elog))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		data := make(glog.JSON)
		for k, v := range c.QueryParams() {
			data[k] = v
		}
		c.Logger().Warn("Try to use query params")
		c.Logger().Printj(data)
		return c.HTML(200, `<a href="/">Home</a><br/><a href="?a=1&b=2&a=3">Try query params</a>`)
	})

	e.Run(standard.New(":8888"))

}
