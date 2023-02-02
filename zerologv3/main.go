package main

import (
	"context"
	"github.com/rs/zerolog"
	loggerI "github.com/tenenwurcel/logger"
	"os"
	"reflect"
)

type logger struct {
	zl zerolog.Logger
}

func (l *logger) Init() {}

func (l *logger) SetLevel(level loggerI.Level) {
	l.zl.Level(zerolog.Level(level))
}

func (l *logger) Debug(msg string, fields interface{}) {
	e := l.zl.Debug()

	makeLog(e, msg, fields)
}

func (l *logger) Info(msg string, fields interface{}) {
	e := l.zl.Info()

	makeLog(e, msg, fields)
}

func (l *logger) Warn(msg string, fields interface{}) {
	e := l.zl.Warn()

	makeLog(e, msg, fields)
}

func (l *logger) Error(msg string, fields interface{}) {
	e := l.zl.Error()

	makeLog(e, msg, fields)
}

func (l *logger) Fatal(msg string, fields interface{}) {
	e := l.zl.Fatal()

	makeLog(e, msg, fields)
}

func (l *logger) Panic(msg string, fields interface{}) {
	e := l.zl.Panic()

	makeLog(e, msg, fields)
}

func makeLog(loggerEvent *zerolog.Event, msg string, fields interface{}) {
	if loggerEvent == nil {
		return
	}

	if fields == nil {
		loggerEvent.Msg(msg)
		return
	}

	val := reflect.ValueOf(fields)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			loggerEvent.Str(val.Type().Field(i).Name, field.String())
		case reflect.Int:
			loggerEvent.Int(val.Type().Field(i).Name, int(field.Int()))
		case reflect.Int8:
			loggerEvent.Int8(val.Type().Field(i).Name, int8(field.Int()))
		case reflect.Int16:
			loggerEvent.Int16(val.Type().Field(i).Name, int16(field.Int()))
		case reflect.Int32:
			loggerEvent.Int32(val.Type().Field(i).Name, int32(field.Int()))
		case reflect.Int64:
			loggerEvent.Int64(val.Type().Field(i).Name, field.Int())
		case reflect.Uint:
			loggerEvent.Uint(val.Type().Field(i).Name, uint(field.Uint()))
		case reflect.Uint8:
			loggerEvent.Uint8(val.Type().Field(i).Name, uint8(field.Uint()))
		case reflect.Uint16:
			loggerEvent.Uint16(val.Type().Field(i).Name, uint16(field.Uint()))
		case reflect.Uint32:
			loggerEvent.Uint32(val.Type().Field(i).Name, uint32(field.Uint()))
		case reflect.Uint64:
			loggerEvent.Uint64(val.Type().Field(i).Name, field.Uint())
		case reflect.Float32:
			loggerEvent.Float32(val.Type().Field(i).Name, float32(field.Float()))
		case reflect.Float64:
			loggerEvent.Float64(val.Type().Field(i).Name, field.Float())
		case reflect.Bool:
			loggerEvent.Bool(val.Type().Field(i).Name, field.Bool())
		case reflect.Interface:
			loggerEvent.Interface(val.Type().Field(i).Name, field)
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

	l2.Debug("Hello world", "a")
}
