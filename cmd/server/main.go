package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"workflow-approval-service/config"
	"workflow-approval-service/handler"
	"workflow-approval-service/repository"
	"workflow-approval-service/router"
	"workflow-approval-service/service"
)

func main() {
	// Load env
	config.LoadEnv()

	// Database
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Run migrations
	if err := config.RunMigration(db); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	// =========================
	// Repositories
	// =========================
	workflowRepo := repository.NewWorkflowRepository()
	workflowStepRepo := repository.NewWorkflowStepRepository()
	requestRepo := repository.NewRequestRepository()

	// =========================
	// Services
	// =========================
	workflowService := service.NewWorkflowService(db, workflowRepo)
	workflowStepService := service.NewWorkflowStepService(db, workflowStepRepo)
	requestService := service.NewRequestService(db, requestRepo, workflowStepRepo)

	// =========================
	// Handlers
	// =========================
	workflowHandler := handler.NewWorkflowHandler(workflowService)
	workflowStepHandler := handler.NewWorkflowStepHandler(workflowStepService)
	requestHandler := handler.NewRequestHandler(requestService)

	// =========================
	// Router
	// =========================
	r := gin.Default()
	router.RegisterWorkflowRoutes(r, workflowHandler)
	router.RegisterWorkflowStepRoutes(r, workflowStepHandler)
	router.RegisterRequestRoutes(r, requestHandler)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
