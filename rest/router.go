package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pararang/hris/rest/handler"
	"github.com/pararang/hris/rest/middleware"
)

// SetupRouter sets up the HTTP router
func setupRouter(
	userHandler *handler.UserHandler,
	attendanceHandler *handler.AttendanceHandler,
	overtimeHandler *handler.OvertimeHandler,
	reimbursementHandler *handler.ReimbursementHandler,
	payslipHandler *handler.PayslipHandler,
	apiKeyMiddleware *middleware.ApiKeyMiddleware,
	authMiddleware *middleware.AuthMiddleware,
	loggerMiddleware *middleware.LoggerMiddleware,
) *gin.Engine {
	router := gin.Default()

	publicV1 := router.Group("/v1")
	{
		publicV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	protectedV1 := router.Group("/v1")
	{
		protectedV1.Use(apiKeyMiddleware.Check())
		protectedV1.POST("/login", userHandler.Login) // Login might also need to be excluded from APIKeyAuth if it's the entry point

		shouldAdmin := protectedV1.Group("")
		shouldAdmin.Use(authMiddleware.AdminAuth())
		shouldAdmin.Use(loggerMiddleware.Logger())
		{
			shouldAdmin.POST("/attendances/periods", attendanceHandler.CreateAttendancePeriod)
			shouldAdmin.POST("/payrolls", payslipHandler.ProcessPayroll)
		}

		/// employee
		attendances := protectedV1.Group("/attendances")
		attendances.Use(authMiddleware.EmployeeAuth())
		{
			attendances.POST("/clockin", attendanceHandler.Clockin)
			attendances.POST("/clockout", attendanceHandler.Clockout)
		}

		overtimes := protectedV1.Group("/overtimes")
		overtimes.Use(authMiddleware.EmployeeAuth())
		{
			overtimes.POST("", overtimeHandler.SubmitOvertime)
		}

		reimbursements := protectedV1.Group("/reimbursements")
		reimbursements.Use(authMiddleware.EmployeeAuth())
		{
			reimbursements.POST("", reimbursementHandler.SubmitReimbursement)
		}
	}

	return router
}
