package logger

import (
	"fmt"
	"os"
	"strings"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel string) *zap.Logger {
	var level zap.AtomicLevel
	var syncOutput zapcore.WriteSyncer

	switch strings.ToLower(logLevel) {
	case "", "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		fmt.Printf("invalid log level '%s' supplied, defaulting to 'info'", logLevel)
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	syncOutput = zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		syncOutput,
		level,
	)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core = zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		syncOutput,
		level,
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return logger
}

func LogWithTraceIdAndSpan(logger *zap.Logger, span trace.Span) *zap.Logger {
	return logger.With(
		zap.String("traceId", span.SpanContext().TraceID().String()),
		zap.String("spanId", span.SpanContext().SpanID().String()),
	)
}
