package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/deepayanMallick/approval-crud/internal/config"
	"github.com/deepayanMallick/approval-crud/internal/db"

	"github.com/pressly/goose/v3"
)

func main() {
	if err := Migrate(); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	const (
		usageRun      = `goose [OPTIONS] COMMAND`
		usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version`
	)
	fmt.Println(usageRun)
	flag.PrintDefaults()
	fmt.Println(usageCommands)
}

// Migrate is a wrapper over goose.Run that read database connections from config file.
func Migrate() error {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	// Initialize database connection
	dbConn, err := db.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("error opening db connection: %v", err)
	}
	defer dbConn.Close()

	// Set dialect
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %v", err)
	}

	// Validate command
	if len(args) == 0 {
		return errors.New("expected at least one arg")
	}

	command := args[0]

	// Get migration directory (default to ./migrations/sql if not specified)
	migrationDir := "./migrations/sql"
	if envDir := os.Getenv("MIGRATION_DIR"); envDir != "" {
		migrationDir = envDir
	}

	// Run goose command
	if err := goose.Run(command, dbConn.DB, migrationDir, args[1:]...); err != nil {
		return fmt.Errorf("goose run: %v", err)
	}

	return nil
}
