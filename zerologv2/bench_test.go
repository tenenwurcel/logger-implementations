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

	s1, _ := loggerI.NewField("service", loggerI.StringType, "dummy")

	s2, _ := loggerI.NewField("version", loggerI.StringType, "0.0.1")

	s3, _ := loggerI.NewField("name", loggerI.StringType, "John Doe")

	i, _ := loggerI.NewField("age", loggerI.IntType, 42)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world", s1, s2, s3, i)
	}
}

func BenchmarkLoggerInterfaceWithZerolog10Fields(b *testing.B) {
	// run the Fib function b.N times
	ml := &logger{
		zl: zerolog.New(io.Discard),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	s1, _ := loggerI.NewField("service", loggerI.StringType, "dummy")

	s2, _ := loggerI.NewField("version", loggerI.StringType, "0.0.1")

	s3, _ := loggerI.NewField("name", loggerI.StringType, "John Doe")

	s4, _ := loggerI.NewField("field1", loggerI.StringType, "value1")

	s5, _ := loggerI.NewField("field2", loggerI.StringType, "value2")

	s6, _ := loggerI.NewField("field3", loggerI.StringType, "value3")

	s7, _ := loggerI.NewField("field4", loggerI.StringType, "value4")

	s8, _ := loggerI.NewField("field5", loggerI.StringType, "value5")

	s9, _ := loggerI.NewField("field6", loggerI.StringType, "value6")

	i, _ := loggerI.NewField("age", loggerI.IntType, 42)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world", s1, s2, s3, s4, s5, s6, s7, s8, s9, i)
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
