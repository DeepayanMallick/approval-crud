package main

import (
	"log"

	"github.com/deepayanMallick/approval-crud/internal/config"
	"github.com/deepayanMallick/approval-crud/internal/db"
	"github.com/deepayanMallick/approval-crud/internal/handlers"
	"github.com/deepayanMallick/approval-crud/internal/repository"
	"github.com/deepayanMallick/approval-crud/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database
	dbConn, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConn.Close()

	// Initialize repository
	approvalRepo := repository.NewApprovalRepository(dbConn)

	// Initialize handlers
	approvalHandler := handlers.NewApprovalHandler(approvalRepo)

	// Initialize router
	router := gin.Default()

	// Setup routes
	routes.SetupApprovalRoutes(router, approvalHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
