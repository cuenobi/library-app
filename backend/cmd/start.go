package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "library-service/docs" // Import docs package
	"library-service/configs"
	"library-service/internal/adapter/jwt"
	"library-service/internal/adapter/postgres"
	"library-service/internal/handler"
	"library-service/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/cobra"
)

// @title Library Service API
// @version 1.0
// @description This is a sample server for a library service.
// @host localhost:8080
// @BasePath /
func start(cmd *cobra.Command, args []string) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))

	ctx := context.Background()
	cfg := configs.GetConfig()

	fmt.Println("App is starting...")

	// Init database
	pg := postgres.NewPostgres(cfg.Postgres, ctx)
	postgres.SeedData(cfg.Postgres, pg)

	// Init fiber server
	f := startServer(cfg.ServerConfig.Port, logger, cfg)

	// Init repositories
	userRepo := postgres.NewUser(pg)
	bookRepo := postgres.NewBook(pg)

	// Init services
	jwtService := jwt.NewJwtToken(cfg.JwtConfig)
	userService := service.NewUserService(userRepo, jwtService)
	bookService := service.NewBookService(bookRepo)

	validate := validator.New()

	// Register routes
	handler.NewRouteUserHandler(f, userService, jwtService, validate)
	handler.NewRouteBookHandler(f, bookService, jwtService, validate)

	// Swagger route
	f.Get("/swagger/*", swagger.HandlerDefault)

	gracefulShutdown(f, logger)

	return nil
}

// startServer initializes a new Fiber application, sets up routes, and starts
// listening on the specified port. It returns the Fiber application instance.
// The provided logger is used for logging server errors, and the postgres connection
// is available for database interactions.
func startServer(port string, slog *slog.Logger, cfg *configs.Config) *fiber.App {
	// Create a new Fiber application
	f := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	f.Use(cors.New())

	f.Use(recover.New(recover.Config{
		EnableStackTrace: cfg.ServerConfig.EnableStackTrace,
	}))

	loggerMiddleware := logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		Format:     "${time} | ${status} | ${latency} | ${ips} | ${method} | ${path}\n",
	})
	f.Use(loggerMiddleware)

	// Define a simple ping route for health checks
	f.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Start the server in a goroutine
	go func() {
		// Listen on the specified port
		if err := f.Listen(":" + port); err != nil {
			// Log any errors encountered while starting the server
			slog.With("error", err).Error("Error starting server")
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
