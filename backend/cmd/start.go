package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"library-service/configs"
	"library-service/internal/adapter/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func start(cmd *cobra.Command, args []string) (err error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))

	ctx := context.Background()

	cfg := configs.GetConfig()

	fmt.Println("App is starting...")

	postgres := postgres.NewPostgres(cfg.Postgres, ctx)

	f := startServer(cfg.ServerConfig.Port, logger, postgres)

	gracefulShutdown(f, logger)

	return nil
}

// startServer initializes a new Fiber application, sets up routes, and starts
// listening on the specified port. It returns the Fiber application instance.
// The provided logger is used for logging server errors, and the postgres connection
// is available for database interactions.
func startServer(port string, logger *slog.Logger, pg *gorm.DB) *fiber.App {
	// Create a new Fiber application
	f := fiber.New()

	// Define a simple ping route for health checks
	f.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Start the server in a goroutine
	go func() {
		// Listen on the specified port
		if err := f.Listen(":" + port); err != nil {
			// Log any errors encountered while starting the server
			logger.Error("Error starting server", slog.String("error", err.Error()))
		}
	}()

	// Return the Fiber application instance
	return f
}

// gracefulShutdown listens for the SIGINT and SIGTERM signals and starts
// the graceful shutdown process once it receives one of these signals.
// It will wait for up to 3 seconds for the server to shut down before
// exiting.
func gracefulShutdown(f *fiber.App, logger *slog.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sigReceived := <-stop
	logger.Info("Received signal", slog.String("signal", sigReceived.String()))

	// The shutdown deadline is set to 3 seconds. If the server does not
	// shut down within this deadline, the program will exit with code 0.
	shutdownDeadline := time.After(3 * time.Second)

	// Shut down the server gracefully.
	if err := f.Shutdown(); err != nil {
		logger.Error("Error shutting down server", slog.Any("error", err))
	}

	// Wait until the server is fully shut down or the deadline is reached.
	select {
	case <-shutdownDeadline:
		logger.Info("Graceful shutdown completed")
	case <-time.After(5 * time.Second):
		logger.Info("Graceful shutdown timed out")
	}

	// Exit the program with code 0.
	os.Exit(0)
}
