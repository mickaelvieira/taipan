package logger

import (
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

// StackTracer --
type StackTracer interface {
	StackTrace() errors.StackTrace
}

// Causer --
type Causer interface {
	Cause() error
}

var l *log.Logger

func getLevel(l string) log.Lvl {
	var a = make(map[string]log.Lvl)
	a["DEBUG"] = log.DEBUG
	a["INFO"] = log.INFO
	a["WARN"] = log.WARN
	a["ERROR"] = log.ERROR
	a["OFF"] = log.OFF
	return a[strings.ToUpper(l)]
}

// Init the logger
func Init(lg *log.Logger, level string) {
	if l != nil {
		panic("Logger has already been initialized")
	}
	l = lg
	l.SetHeader("${time_rfc3339} ${level}")
	l.SetLevel(getLevel(level))
}

// Debug prints stuff in dev mode only
func Debug(i interface{}) {
	l.Debug(i)
}

// Info prints stuff in dev mode only
func Info(i interface{}) {
	l.Info(i)
}

// Warn prints warning messages
func Warn(i interface{}) {
	l.Warn(i)
}

// Error prints error messages
func Error(i interface{}) {
	if err, ok := i.(Causer); ok {
		l.Error(fmt.Sprintf("%+v", err))
	} else {
		l.Error(i)
	}
}

// Fatal prints fatal messages
func Fatal(i interface{}) {
	l.Fatal(i)
}
