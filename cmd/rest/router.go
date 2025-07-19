package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/handler"
	"github.com/prrng/dealls/interface/api/middleware"
)

// SetupRouter sets up the HTTP router
func setupRouter(
	userHandler *handler.UserHandler,
	attendanceHandler *handler.AttendanceHandler,
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

	// Admin routes
	admin := router.Group("/v1/admin")
	admin.Use(authMiddleware.AdminAuth())
	admin.Use(loggerMiddleware.Logger())
	{
		admin.POST("/payroll-periods", attendanceHandler.CreateAttendancePeriod)

	}

	// Employee routes
	employee := router.Group("/api/employee")
	employee.Use(authMiddleware.EmployeeAuth())
	employee.Use(loggerMiddleware.Logger())
	{
		// employee.GET("/profile", userHandler.GetEmployeeProfile)
		// employee.PUT("/profile", userHandler.UpdateEmployeeProfile)
	}

	return router
}
