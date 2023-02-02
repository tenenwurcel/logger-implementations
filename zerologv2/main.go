package main

import (
	"context"
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

func (l *logger) Debug(msg string, fields ...loggerI.Field) {
	e := l.zl.Debug()

	makeLog(e, msg, fields...)
}

func (l *logger) Info(msg string, fields ...loggerI.Field) {
	e := l.zl.Info()

	makeLog(e, msg, fields...)
}

func (l *logger) Warn(msg string, fields ...loggerI.Field) {
	e := l.zl.Warn()

	makeLog(e, msg, fields...)
}

func (l *logger) Error(msg string, fields ...loggerI.Field) {
	e := l.zl.Error()

	makeLog(e, msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...loggerI.Field) {
	e := l.zl.Fatal()

	makeLog(e, msg, fields...)
}

func (l *logger) Panic(msg string, fields ...loggerI.Field) {
	e := l.zl.Panic()

	makeLog(e, msg, fields...)
}

func makeLog(loggerEvent *zerolog.Event, msg string, fields ...loggerI.Field) {
	if loggerEvent == nil {
		return
	}

	for _, f := range fields {
		switch f.GetType() {
		case loggerI.StringType:
			v := f.GetValue().(string)
			loggerEvent = loggerEvent.Str(f.GetKey(), v)
		case loggerI.IntType:
			v := f.GetValue().(int)
			loggerEvent = loggerEvent.Int(f.GetKey(), v)
		case loggerI.Int64Type:
			v := f.GetValue().(int64)
			loggerEvent = loggerEvent.Int64(f.GetKey(), v)
		case loggerI.Float64Type:
			v := f.GetValue().(float64)
			loggerEvent = loggerEvent.Float64(f.GetKey(), v)
		case loggerI.BoolType:
			v := f.GetValue().(bool)
			loggerEvent = loggerEvent.Bool(f.GetKey(), v)
		case loggerI.InterfaceType:
			loggerEvent = loggerEvent.Interface(f.GetKey(), f.GetValue())
		}
	}

	loggerEvent.Msg(msg)
}

func main() {
	ml := &logger{
		zl: zerolog.New(os.Stdout),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	ctx := context.Background()

	ctx = l.WithContext(ctx)

	l2 := loggerI.FromContext(ctx)

	l2.Debug("Hello world", nil)
}

/*switch f.GetType() {
case loggerI.StringType:
	loggerEvent = loggerEvent.Str(f.GetKey(), f.ToString())
case loggerI.IntType:
	loggerEvent = loggerEvent.Int(f.GetKey(), f.ToInt())
case loggerI.Int64Type:
	loggerEvent = loggerEvent.Int64(f.GetKey(), f.ToInt64())
case loggerI.Float64Type:
	loggerEvent = loggerEvent.Float64(f.GetKey(), f.ToFloat64())
case loggerI.BoolType:
	loggerEvent = loggerEvent.Bool(f.GetKey(), f.ToBool())
case loggerI.InterfaceType:
	loggerEvent = loggerEvent.Interface(f.GetKey(), f.ToInterface())
}*/
