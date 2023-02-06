package main

import (
	loggerI "github.com/tenenwurcel/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"testing"
)

func BenchmarkLoggerInterfaceWithZap(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.DebugLevel,
	)
	zapLogger := zap.New(core)

	ml := &logger{
		zl: zapLogger,
	}

	l := loggerI.New(ml)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world",
			"name", "John Doe",
			"age", 42,
			"service", "dummy",
			"version", "0.0.1",
		)
	}
}

func BenchmarkLoggerInterfaceWithZap10Fields(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.DebugLevel,
	)
	zapLogger := zap.New(core)

	ml := &logger{
		zl: zapLogger,
	}

	l := loggerI.New(ml)

	l.SetLevel(loggerI.DebugLevel)

	for n := 0; n < b.N; n++ {
		l.Debug("Hello world",
			"name", "John Doe",
			"age", 42,
			"service", "dummy",
			"version", "0.0.1",
			"field1", "value1",
			"field2", "value2",
			"field3", "value3",
			"field4", "value4",
			"field5", "value5",
			"field6", "value6",
		)
	}
}

func BenchmarkZap(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.DebugLevel,
	)
	zl := zap.New(core)

	for n := 0; n < b.N; n++ {
		zl.Debug(
			"Hello world",
			zap.String("service", "dummy"),
			zap.String("version", "0.0.1"),
			zap.String("name", "John Doe"),
			zap.Int("age", 42),
		)
	}
}

func BenchmarkZap10Fields(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.DebugLevel,
	)
	zl := zap.New(core)

	for n := 0; n < b.N; n++ {
		zl.Debug(
			"Hello world",
			zap.String("service", "dummy"),
			zap.String("version", "0.0.1"),
			zap.String("name", "John Doe"),
			zap.Int("age", 42),
			zap.String("field1", "value1"),
			zap.String("field2", "value2"),
			zap.String("field3", "value3"),
			zap.String("field4", "value4"),
			zap.String("field5", "value5"),
			zap.String("field6", "value6"),
		)
	}
}
