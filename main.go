package main

import (
	"log/slog"
	"runtime"
	"time"

	"github.com/brody192/logger"
)

func main() {
	ticker := time.NewTicker(1200 * time.Millisecond)
	defer ticker.Stop()

	extraAttrs := []any{
		slog.String("go_version", runtime.Version()),
		slog.String("test_string", "Hello, World!"),
		slog.Int64("test_int", 12345678),
		slog.Float64("test_float", 1.2345678),
		slog.Any("test_string_slice", []string{"hello", "world"}),
		slog.Any("test_int_slice", []int{123, 456, 789}),
		slog.Group("test_group",
			slog.String("test_grouped_string", "Hello, World!"),
			slog.Int64("test_grouped_int", 12345678),
		),
	}

	loggerStdout := logger.Stdout.With(extraAttrs...)
	// loggerStderr := logger.Stderr.With(extraAttrs...)

	// var counter int64 = 0

	var printLogs = func() {
		// fmt.Println(counter)
		// counter++
		loggerStdout.Info("INFO level -> stdout")
		loggerStdout.Warn("WARN level -> stdout")
		loggerStdout.Error("ERROR level -> stdout")
		loggerStdout.Debug("DEBUG level -> stdout")

		// loggerStderr.Info("INFO level -> stderr")
		// loggerStderr.Warn("WARN level -> stderr")
		// loggerStderr.Error("ERROR level -> stderr")
	}

	printLogs()

	for range ticker.C {
		printLogs()
	}
}
