package main

import (
	log "github.com/sirupsen/logrus"
	loggerI "github.com/tenenwurcel/logger"
	"io"
	"testing"
)

func BenchmarkLogrusWithInterface(b *testing.B) {
	custom := log.New()
	custom.SetOutput(io.Discard)
	custom.SetFormatter(&log.JSONFormatter{})

	ml := &logger{
		logrus: custom,
	}

	l := loggerI.New(ml)
	l.SetLevel(loggerI.DebugLevel)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello World",
			"name", "John Doe",
			"age", 42,
			"app_name", "dummy",
			"app_version", "0.0.1",
		)
	}
}

func BenchmarkLogrus10FieldsWithInterface(b *testing.B) {
	custom := log.New()
	custom.SetOutput(io.Discard)
	custom.SetFormatter(&log.JSONFormatter{})

	ml := &logger{
		logrus: custom,
	}

	l := loggerI.New(ml)
	l.SetLevel(loggerI.DebugLevel)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello Worlds",
			"name", "John Doe",
			"age", 42,
			"app_name", "dummy",
			"app_version", "0.0.1",
			"field1", "value1",
			"field2", "value2",
			"field3", "value3",
			"field4", "value4",
			"field5", "value5",
			"field6", "value6",
		)
	}
}

func BenchmarkLogrus(b *testing.B) {
	custom := log.New()
	custom.SetOutput(io.Discard)
	custom.SetLevel(log.DebugLevel)
	custom.SetFormatter(&log.JSONFormatter{})

	for n := 0; n < b.N; n++ {
		custom.WithFields(log.Fields{
			"name":        "John Doe",
			"age":         42,
			"app_name":    "dummy",
			"app_version": "0.0.1",
		}).Debug("Hello world")
	}
}

func BenchmarkLogrus10Fields(b *testing.B) {
	custom := log.New()
	custom.SetOutput(io.Discard)
	custom.SetLevel(log.DebugLevel)
	custom.SetFormatter(&log.JSONFormatter{})

	for n := 0; n < b.N; n++ {
		custom.WithFields(log.Fields{
			"name":        "John Doe",
			"age":         42,
			"app_name":    "dummy",
			"app_version": "0.0.1",
			"field1":      "value1",
			"field2":      "value2",
			"field3":      "value3",
			"field4":      "value4",
			"field5":      "value5",
			"field6":      "value6",
		}).Debug("Hello world")
	}
}
