package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"testing"
)

func BenchmarkLogrus(b *testing.B) {
	custom := log.New()

	// Log as JSON instead of the default ASCII formatter.
	//custom.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	custom.SetOutput(io.Discard)

	// Only log the warning severity or above.
	custom.SetLevel(log.DebugLevel)

	for n := 0; n < b.N; n++ {
		custom.WithFields(log.Fields{
			"name":    "John Doe",
			"age":     42,
			"service": "dummy",
			"version": "0.0.1",
		}).Debug("Hello world")
	}
}

func BenchmarkLogrus10Fields(b *testing.B) {
	custom := log.New()

	// Log as JSON instead of the default ASCII formatter.
	//custom.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	custom.SetOutput(io.Discard)

	// Only log the warning severity or above.
	custom.SetLevel(log.DebugLevel)

	for n := 0; n < b.N; n++ {
		custom.WithFields(log.Fields{
			"name":    "John Doe",
			"age":     42,
			"service": "dummy",
			"version": "0.0.1",
			"field1":  "value1",
			"field2":  "value2",
			"field3":  "value3",
			"field4":  "value4",
			"field5":  "value5",
			"field6":  "value6",
		}).Debug("Hello world")
	}
}
