package main

import (
	loggerI "github.com/tenenwurcel/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type logger struct {
	zl *zap.Logger
}

func (l *logger) Init() {}

func (l *logger) SetLevel(level loggerI.Level) {}

func (l *logger) Debug(msg string, fields ...interface{}) {
	zapFields := makeFields(fields)

	l.zl.Debug(msg, zapFields...)
}

func (l *logger) Info(msg string, fields ...interface{}) {
	zapFields := makeFields(fields)

	l.zl.Info(msg, zapFields...)
}

func (l *logger) Warn(msg string, fields ...interface{}) {
	zapFields := makeFields(fields)

	l.zl.Warn(msg, zapFields...)
}

func (l *logger) Error(msg string, fields ...interface{}) {
	zapFields := makeFields(fields)

	l.zl.Error(msg, zapFields...)
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	zapFields := makeFields(fields)

	l.zl.Fatal(msg, zapFields...)
}

func (l *logger) Panic(msg string, fields ...interface{}) {
	zapFields := makeFields(fields)

	l.zl.Panic(msg, zapFields...)
}

func makeFields(fields []interface{}) []zap.Field {
	var zapFields []zap.Field

	for i := 0; i < len(fields); i = i + 2 {
		key := fields[i].(string)

		zapField := zap.Any(key, fields[i])

		zapFields = append(zapFields, zapField)
	}

	return zapFields
}

func main() {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	zapLogger := zap.New(core)

	ml := &logger{
		zl: zapLogger,
	}

	l := loggerI.New(ml)
	l.Debug("Hello world", "key", "value", "keyInt", 42)
}
