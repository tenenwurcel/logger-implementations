package main

import (
	"github.com/rs/zerolog"
	loggerI "github.com/tenenwurcel/logger"
	"io"
	"testing"
)

func BenchmarkLoggerInterfaceWithZerolog(b *testing.B) {
	// run the Fib function b.N times
	ml := &logger{
		zl: zerolog.New(io.Discard),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world",
			loggerI.String("service", "dummy"),
			loggerI.String("version", "0.0.1"),
			loggerI.String("name", "John Doe"),
			loggerI.Int("age", 42),
		)
	}
}

func BenchmarkLoggerInterfaceWithZerolog10Fields(b *testing.B) {
	// run the Fib function b.N times
	ml := &logger{
		zl: zerolog.New(io.Discard),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world",
			loggerI.String("field1", "value1"),
			loggerI.String("field2", "value2"),
			loggerI.String("field3", "value3"),
			loggerI.String("field4", "value4"),
			loggerI.String("field5", "value5"),
			loggerI.String("field6", "value6"),
			loggerI.String("service", "dummy"),
			loggerI.String("version", "0.0.1"),
			loggerI.String("name", "John Doe"),
			loggerI.Int("age", 42),
		)
	}
}

func BenchmarkZerolog(b *testing.B) {
	zl := zerolog.New(io.Discard)

	for n := 0; n < b.N; n++ {
		zl.Debug().
			Str("service", "dummy").
			Str("version", "0.0.1").
			Str("name", "John Doe").
			Int("age", 42).
			Msg("Hello world")
	}
}

func BenchmarkZerolog10Fields(b *testing.B) {
	zl := zerolog.New(io.Discard)

	for n := 0; n < b.N; n++ {
		zl.Debug().
			Str("field1", "value1").
			Str("field2", "value2").
			Str("field3", "value3").
			Str("field4", "value4").
			Str("field5", "value5").
			Str("field6", "value6").
			Str("service", "dummy").
			Str("version", "0.0.1").
			Str("name", "John Doe").
			Int("age", 42).
			Msg("Hello world")
	}
}
