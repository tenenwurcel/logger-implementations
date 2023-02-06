package main

import (
	log "github.com/sirupsen/logrus"
	loggerI "github.com/tenenwurcel/logger"
	"os"
)

type logger struct {
	logrus *log.Logger
}

func (l *logger) Init() {}

func (l *logger) SetLevel(level loggerI.Level) {
	logrusLevelCalc := 5 - level
	logrusLevel := log.Level(logrusLevelCalc)

	l.logrus.SetLevel(logrusLevel)
}

func (l *logger) Debug(msg string, fields ...interface{}) {
	logrusFields := makeFields(fields)

	l.logrus.
		WithFields(logrusFields).
		Debug(msg)

}

func (l *logger) Info(msg string, fields ...interface{}) {
	logrusFields := makeFields(fields)

	l.logrus.
		WithFields(logrusFields).
		Info(msg)
}

func (l *logger) Warn(msg string, fields ...interface{}) {
	logrusFields := makeFields(fields)

	l.logrus.
		WithFields(logrusFields).
		Warn(msg)
}

func (l *logger) Error(msg string, fields ...interface{}) {
	logrusFields := makeFields(fields)

	l.logrus.
		WithFields(logrusFields).
		Error(msg)
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	logrusFields := makeFields(fields)

	l.logrus.
		WithFields(logrusFields).
		Fatal(msg)
}

func (l *logger) Panic(msg string, fields ...interface{}) {
	logrusFields := makeFields(fields)

	l.logrus.
		WithFields(logrusFields).
		Panic(msg)
}

func makeFields(fields []interface{}) log.Fields {
	logrusFields := log.Fields{}

	for i := 0; i < len(fields); i = i + 2 {
		key := (fields[i]).(string)

		logrusFields[key] = fields[i+1]
	}

	return logrusFields
}

func main() {
	custom := log.New()
	custom.SetOutput(os.Stdout)
	custom.SetFormatter(&log.JSONFormatter{})

	ml := &logger{
		logrus: custom,
	}

	l := loggerI.New(ml)
	l.SetLevel(loggerI.DebugLevel)

	l.Debug("debug", "key", "value", "keyInt", 42)
}
