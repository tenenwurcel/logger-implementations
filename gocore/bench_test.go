package main

import (
	"github.com/dock-tech/dock-gocore-lib/log"
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
