package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/internals/events"
)

func main() {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		logger.Warn("Could not load .env file, using system environment variables")
	}

	// Ensure required environment variables are set
	requiredEnvVars := []string{"GOOGLE_PUBSUB_SUBSCRIPTION"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			slog.Error("Missing required environment variable", "envVar", envVar)
			os.Exit(1)
		}
	}

	// Gracefully handle shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Start the listener
	go func() {
		logger.Info("Starting Pub/Sub listener")
		if err := events.ListenAndDispatch(); err != nil {
			logger.Error("Listener encountered an error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for shutdown signal
	<-stopChan
	logger.Info("Shutting down listener")
}
