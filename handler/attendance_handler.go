package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/dto"
	"github.com/prrng/dealls/libs/auth"
	"github.com/prrng/dealls/libs/httpresp"
	"github.com/prrng/dealls/libs/logger"
)

// AttendanceHandler handles attendance-related HTTP requests
type AttendanceHandler struct {
	attendanceUseCase usecase.AttendanceUseCase
	log               logger.Logger
}

// NewAttendanceHandler creates a new instance of AttendanceHandler
func NewAttendanceHandler(attendanceUseCase usecase.AttendanceUseCase, log logger.Logger) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceUseCase: attendanceUseCase,
		log:               log,
	}
}

func (h *AttendanceHandler) CreateAttendancePeriod(c *gin.Context) {
	var req dto.CreateAttendancePeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	createdBy, ok := c.Get(auth.CtxKeyAuthUserEmail)
	if !ok {
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unauthorized user")))
		return
	}

	dateParser := func(date string) (time.Time, error) {
		return time.Parse(time.DateOnly, date)
	}

	startDate, err := dateParser(req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	endDate, err := dateParser(req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	if startDate.Compare(endDate) != -1 {
		c.JSON(http.StatusBadRequest, httpresp.Err(errors.New("start date must be before end date")))
		return
	}

	period := &entity.PayrollPeriod{
		StartDate: startDate,
		EndDate:   endDate,
		Status:    entity.PayrollPeriodStatusPending,
		BaseModel: entity.BaseModel{
			CreatedBy: createdBy.(string),
		},
	}

	createdPeriod, err := h.attendanceUseCase.CreateAttendancePeriod(c.Request.Context(), period)
	if err != nil {
		h.log.Warn(err.Error())
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(dto.AttendancePeriodResponse{
		ID:        createdPeriod.ID,
		StartDate: createdPeriod.StartDate,
		EndDate:   createdPeriod.EndDate,
		Status:    createdPeriod.Status,
	}))
}
