// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"syscall"

	"log/slog"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/appclient"
	"github.com/CamPlume1/khoury-classroom/internal/server"
	"github.com/CamPlume1/khoury-classroom/internal/storage/postgres"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	// Load environment variables if running locally
	if isLocal() {
		if err := godotenv.Load("../../../.env"); err != nil {
			log.Fatalf("Unable to load environment variables necessary for application: %v", err)
		}
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		log.Fatalf("Unable to load configuration: %v", err)
	}
	log.Println("Configuration loaded:", cfg)

	// Initialize the database connection pool
	db, err := postgres.New(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("Failed to establish database connection: %v", err)
	}
	// Ensure the connection pool is closed when main exits
	defer db.Close(context.Background())

	// Run database migrations
	if err := runMigrations(ctx, db, "./database/migration"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations applied successfully!")

	// Initialize GitHub App Client
	GitHubApp, err := appclient.New(&cfg.GitHubAppClient)
	if err != nil {
		log.Fatalf("Unable to establish connection with GitHub: %v", err)
	}

	// Initialize the server
	app := server.New(types.Params{
		Store:     db,
		GitHubApp: GitHubApp,
		UserCfg:   cfg.GitHubUserClient,
	})

	// Start the server in a separate goroutine
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Setup channel to listen for OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	// Begin shutdown process
	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("Failed to shutdown server", "error", err)
	}

	slog.Info("Server shutdown complete")
}

// runMigrations loads and executes all SQL migration files from the specified directory.
func runMigrations(ctx context.Context, db *postgres.DB, migrationsDir string) error {
	// Read all files from the migrations directory
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory '%s': %w", migrationsDir, err)
	}

	// Filter and collect only .sql files
	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}

	// Sort the migration files to ensure correct execution order
	sort.Strings(migrationFiles)

	// Execute each migration file sequentially
	for _, fileName := range migrationFiles {
		filePath := filepath.Join(migrationsDir, fileName)
		log.Printf("Applying migration: %s", fileName)

		// Read the migration file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file '%s': %w", fileName, err)
		}

		// Execute the migration SQL
		_, err = db.Exec(ctx, string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration '%s': %w", fileName, err)
		}

		log.Printf("Successfully applied migration: %s", fileName)
	}

	return nil
}

// isLocal determines if the application is running in a local environment.
func isLocal() bool {
	return os.Getenv("APP_ENVIRONMENT") == "LOCAL"
}
