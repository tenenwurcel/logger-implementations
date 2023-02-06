package main

import (
	"github.com/rs/zerolog"
	loggerI "github.com/tenenwurcel/logger"
	"os"
)

type logger struct {
	zl zerolog.Logger
}

func (l *logger) Init() {}

func (l *logger) SetLevel(level loggerI.Level) {
	l.zl.Level(zerolog.Level(level))
}

func (l *logger) Debug(msg string, fields ...interface{}) {
	e := l.zl.Debug()

	makeLog(e, msg, fields)
}

func (l *logger) Info(msg string, fields ...interface{}) {
	e := l.zl.Info()

	makeLog(e, msg, fields)
}

func (l *logger) Warn(msg string, fields ...interface{}) {
	e := l.zl.Warn()

	makeLog(e, msg, fields)
}

func (l *logger) Error(msg string, fields ...interface{}) {
	e := l.zl.Error()

	makeLog(e, msg, fields)
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	e := l.zl.Fatal()

	makeLog(e, msg, fields)
}

func (l *logger) Panic(msg string, fields ...interface{}) {
	e := l.zl.Panic()

	makeLog(e, msg, fields)
}

func makeLog(loggerEvent *zerolog.Event, msg string, fields []interface{}) {
	loggerEvent.Fields(fields)
	loggerEvent.Msg(msg)
}

func main() {
	ml := &logger{
		zl: zerolog.New(os.Stdout),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)
	l.Debug("Hello world", "key", "value", "keyInt", 42)
}
