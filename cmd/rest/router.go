package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/handler"
	"github.com/prrng/dealls/interface/api/middleware"
)

// SetupRouter sets up the HTTP router
func setupRouter(
	userHandler *handler.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
	loggerMiddleware *middleware.LoggerMiddleware,
) *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/v1")
	{
		// Authentication
		public.POST("/login", userHandler.Login)

		// Health check
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	// Employee routes
	employee := router.Group("/api/employee")
	employee.Use(authMiddleware.EmployeeAuth())
	employee.Use(loggerMiddleware.Logger())
	{
		// employee.GET("/profile", userHandler.GetEmployeeProfile)
		// employee.PUT("/profile", userHandler.UpdateEmployeeProfile)
	}

	// Admin routes
	admin := router.Group("/api/admin")
	admin.Use(authMiddleware.AdminAuth())
	admin.Use(loggerMiddleware.Logger())
	{
		// User management
		// admin.POST("/employee", userHandler.RegisterEmployee)
		// admin.GET("/employee", userHandler.ListEmployees)

	}

	return router
}
