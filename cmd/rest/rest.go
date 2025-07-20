package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prrng/dealls/config"
	"github.com/prrng/dealls/dbase"
	"github.com/prrng/dealls/handler"
	"github.com/prrng/dealls/interface/api/middleware"
	"github.com/prrng/dealls/libs/auth"
	"github.com/prrng/dealls/libs/logger"
	"github.com/prrng/dealls/repository"
	"github.com/prrng/dealls/usecase"
)

func Serve() {
	cfg := config.New()

	db, err := dbase.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	attendanceRepo := repository.NewAttendanceRepository(db)
	overtimeRepo := repository.NewOvertimeRepository(db)
	reimbursementRepo := repository.NewReimbursementRepository(db)
	auditRepo := repository.NewAuditRepository(db)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo, auditRepo)
	attendanceUseCase := usecase.NewAttendanceUseCase(attendanceRepo)
	overtimeUseCase := usecase.NewOvertimeUseCase(overtimeRepo, attendanceRepo)
	reimbursementUseCase := usecase.NewReimbursementUseCase(reimbursementRepo, attendanceRepo)
	auditUseCase := usecase.NewAuditUseCase(auditRepo)

	jwtService := auth.NewJWTService(cfg.JWT.Secret, time.Duration(cfg.JWT.ExpiryMinutes)*time.Minute)
	libLog := logger.New()

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase, jwtService, libLog)
	attendanceHandler := handler.NewAttendanceHandler(attendanceUseCase, libLog)
	overtimeHandler := handler.NewOvertimeHandler(overtimeUseCase)
	reimbursementHandler := handler.NewReimbursementHandler(reimbursementUseCase)

	authMiddleware := middleware.NewAuthMiddleware(jwtService)
	loggerMiddleware := middleware.NewLoggerMiddleware(auditUseCase)

	router := setupRouter(
		userHandler,
		attendanceHandler,
		overtimeHandler,
		reimbursementHandler,
		authMiddleware,
		loggerMiddleware,
	)

	// Setup server with graceful shutdown
	srv := &http.Server{
		Addr:    ":" + fmt.Sprintf("%d", cfg.Port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %d...", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
