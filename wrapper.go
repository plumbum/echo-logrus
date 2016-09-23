package echologrus

import (
	"io"
	"github.com/labstack/gommon/log"
	"github.com/Sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
	jsonMsg string
}

func NewStandart() *Logger {
	lg := new(Logger)
	lg.Logger = logrus.StandardLogger()
	return lg
}

func New(l *logrus.Logger) *Logger {
	lg := new(Logger)
	lg.Logger = l
	return lg
}

func (l *Logger) SetJsonMsg(msg string) {
	l.jsonMsg = msg
}

func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.Out = w
}

func (l *Logger) SetLevel(lvl log.Lvl) {

	var logrusLvl logrus.Level
	switch lvl {
	case log.DEBUG:
		logrusLvl = logrus.DebugLevel
	case log.INFO:
		logrusLvl = logrus.InfoLevel
	case log.WARN:
		logrusLvl = logrus.WarnLevel
	case log.ERROR:
		logrusLvl = logrus.ErrorLevel
	case log.FATAL:
		logrusLvl = logrus.FatalLevel
	case log.OFF:
		logrusLvl = logrus.DebugLevel
	}
	l.Logger.Level = logrusLvl
}

// func (l *Logger) Print(...interface{}) { }
// func (l *Logger) Printf(string, ...interface{}) { }
func (l *Logger) Printj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Print(l.jsonMsg)
}

// func (l *Logger) Debug(...interface{}) { }
// func (l *Logger) Debugf(string, ...interface{}) { }
func (l *Logger) Debugj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Debug(l.jsonMsg)
}

// func (l *Logger) Info(...interface{}) { }
// func (l *Logger) Infof(string, ...interface{}) { }
func (l *Logger) Infoj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Info(l.jsonMsg)
}

// func (l *Logger) Warn(...interface{}) { }
// func (l *Logger) Warnf(string, ...interface{}) { }
func (l *Logger) Warnj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Warn(l.jsonMsg)
}

// func (l *Logger) Error(...interface{}) { }
// func (l *Logger) Errorf(string, ...interface{}) { }
func (l *Logger) Errorj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Error(l.jsonMsg)
}

// func (l *Logger) Fatal(...interface{}) { }
// func (l *Logger) Fatalf(string, ...interface{}) { }
func (l *Logger) Fatalj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Fatal(l.jsonMsg)
}

func (l *Logger) Panicj(json log.JSON) {
	l.Logger.WithFields(logrus.Fields(json)).Panic(l.jsonMsg)
}

