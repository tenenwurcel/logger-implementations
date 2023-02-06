package main

import (
	"github.com/dock-tech/dock-gocore-lib/log"
	"github.com/dock-tech/dock-gocore-lib/logger"
	"io"
	"testing"
)

func BenchmarkGocoreWith10Fields(b *testing.B) {
	log.SetOutput(io.Discard)
	log.SetLevel(log.DebugLevel)
	log.InitializeLogs(log.Fields{"app_name": "dummy", "app_version": "0.0.1"})

	for n := 0; n < b.N; n++ {
		log.Debugj(Fields{
			Message: "Hello Worlds",
			Name:    "John Doe",
			Age:     42,
			Field1:  "value1",
			Field2:  "value2",
		})
	}
}

func BenchmarkGocoreWith10FieldsWithLoggerConfig(b *testing.B) {
	appName := "dummy"
	appVersion := "0.0.1"
	env := "DEV"

	log.SetLevel(log.DebugLevel)
	log.InitializeLogs(log.Fields{"app_name": appName, "app_version": appVersion})
	logger.Config("panic: ", appName, appVersion, env)
	log.SetOutput(io.Discard)

	for n := 0; n < b.N; n++ {
		log.Debugj(Fields{
			Message: "Hello Worlds",
			Name:    "John Doe",
			Age:     42,
			Field1:  "value1",
		})
	}
}
