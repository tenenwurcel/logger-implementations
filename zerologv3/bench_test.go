package main

import (
	"github.com/rs/zerolog"
	loggerI "github.com/tenenwurcel/logger"
	"io"
	"testing"
)

type Case1 struct {
	Name    string
	Age     int
	Version string
	Service string
}

type Case2 struct {
	Name    string
	Age     int
	Version string
	Service string
	Field1  string
	Field2  string
	Field3  string
	Field4  string
	Field5  string
	Field6  string
}

func BenchmarkLoggerInterfaceWithZerolog(b *testing.B) {
	// run the Fib function b.N times
	ml := &logger{
		zl: zerolog.New(io.Discard),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	c1 := &Case1{
		Name:    "John Doe",
		Age:     42,
		Version: "0.0.1",
		Service: "dummy",
	}

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world", c1)
	}
}

func BenchmarkLoggerInterfaceWithZerolog10Fields(b *testing.B) {
	// run the Fib function b.N times
	ml := &logger{
		zl: zerolog.New(io.Discard),
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	c2 := &Case2{
		Name:    "John Doe",
		Age:     42,
		Version: "0.0.1",
		Service: "dummy",
		Field1:  "value1",
		Field2:  "value2",
		Field3:  "value3",
		Field4:  "value4",
		Field5:  "value5",
		Field6:  "value6",
	}

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world", c2)
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
