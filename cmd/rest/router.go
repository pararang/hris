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
	overtimeHandler *handler.OvertimeHandler,
	reimbursementHandler *handler.ReimbursementHandler,
	authMiddleware *middleware.AuthMiddleware,
	loggerMiddleware *middleware.LoggerMiddleware,
) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", userHandler.Login)
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		admin := v1.Group("")
		admin.Use(authMiddleware.AdminAuth())
		admin.Use(loggerMiddleware.Logger())
		{
			admin.POST("/attendances/periods", attendanceHandler.CreateAttendancePeriod)
		}

		attendances := v1.Group("/attendances")
		attendances.Use(authMiddleware.EmployeeAuth())
		{
			attendances.POST("/clockin", attendanceHandler.Clockin)
			attendances.POST("/clockout", attendanceHandler.Clockout)
		}

		overtimes := v1.Group("/overtimes")
		overtimes.Use(authMiddleware.EmployeeAuth())
		{
			overtimes.POST("", overtimeHandler.SubmitOvertime)
		}

		reimbursements := v1.Group("/reimbursements")
		reimbursements.Use(authMiddleware.EmployeeAuth())
		{
			reimbursements.POST("", reimbursementHandler.SubmitReimbursement)
		}
	}

	return router
}
