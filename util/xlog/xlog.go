package xlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

// InitXLog
//
// Init XLog.
func InitXLog() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

// Info
//
// Output info level log.
func Info(ctx context.Context, msg string) {
	slog.Info(msg, getSource(), getTraceId(ctx))
}

// Debug
//
// Output debug level log.
func Debug(ctx context.Context, msg string) {
	slog.Debug(msg, getSource(), getTraceId(ctx))
}

// Warn
//
// Output warn level log.
func Warn(ctx context.Context, msg string) {
	slog.Warn(msg, getSource(), getTraceId(ctx))
}

// Error
//
// Output error level log.
func Error(ctx context.Context, err error, msg string) {
	slog.Error(msg, getSource(), getTraceId(ctx), slog.String("error", err.Error()))
}

// Fatal
//
// Output fatal level log and exit application.
func Fatal(ctx context.Context, err error, msg string) {
	Error(ctx, err, msg)
	os.Exit(1)
}

// getSource
//
// Get log source.
func getSource() slog.Attr {
	_, file, line, ok := runtime.Caller(2)
	source := "unknown source"
	if ok {
		source = fmt.Sprintf("%s:%d", file, line)
	}
	return slog.String("source", source)
}

// getTraceId
//
// Get traceId from context.
func getTraceId(ctx context.Context) slog.Attr {
	var traceId string
	if ctx != nil && ctx.Value("traceId") != nil {
		traceId = ctx.Value("traceId").(string)
	}
	if traceId == "" {
		traceId = "no traceId"
	}
	return slog.String("traceId", traceId)
}
