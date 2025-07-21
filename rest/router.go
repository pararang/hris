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
		protectedV1.POST("/login", userHandler.Login)

		shouldAdmin := protectedV1.Group("")
		shouldAdmin.Use(authMiddleware.AdminAuth())
		shouldAdmin.Use(loggerMiddleware.Logger())
		{
			shouldAdmin.POST("/attendances/periods", attendanceHandler.CreateAttendancePeriod)
			shouldAdmin.POST("/payrolls", payslipHandler.ProcessPayroll)
		}

		attendances := protectedV1.Group("/attendances")
		attendances.Use(authMiddleware.EmployeeAuth())
		{
			attendances.POST("/clockin", attendanceHandler.Clockin)
			attendances.POST("/clockout", attendanceHandler.Clockout)
		}

		protectedV1.POST("/overtimes", authMiddleware.EmployeeAuth(), overtimeHandler.SubmitOvertime)
		protectedV1.POST("/reimbursements", authMiddleware.EmployeeAuth(), reimbursementHandler.SubmitReimbursement)
		protectedV1.GET("payslips", authMiddleware.EmployeeAuth(), payslipHandler.ListUserPayslips)
		protectedV1.GET("payslips/:id", authMiddleware.EmployeeAuth(), payslipHandler.GetPayslipDetails)
	}

	return router
}
