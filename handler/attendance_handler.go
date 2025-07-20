package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/dto"
	"github.com/prrng/dealls/libs"
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
	ctx := c.Request.Context()
	var req dto.CreateAttendancePeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
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

	createdPeriod, err := h.attendanceUseCase.CreateAttendancePeriod(ctx, dto.CreateAttendancePeriodParam{
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		h.log.Warn(err.Error())
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(dto.AttendancePeriodResponse{
		ID:        createdPeriod.ID,
		StartDate: createdPeriod.StartDate.Format(time.DateOnly),
		EndDate:   createdPeriod.EndDate.Format(time.DateOnly),
		Status:    createdPeriod.Status,
	}))
}

func (h *AttendanceHandler) Clockin(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := ctx.Value(auth.CtxKeyAuthUserID).(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unauthorized user")))
		return
	}

	createdAttendance, err := h.attendanceUseCase.ClockIn(ctx, userID)
	if err != nil {
		switch true {
		case errors.Is(err, libs.ErrWeekendNotAllowed{}):
			c.JSON(http.StatusUnprocessableEntity, httpresp.Err(err))
		default:
			c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		}

		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(dto.ClockinResponse{
		ID:         createdAttendance.ID,
		Date:       createdAttendance.Date.Format(time.DateOnly),
		ClockinAt:  createdAttendance.ClockinAt,
		ClockoutAt: createdAttendance.ClockoutAt,
	}))
}

func (h *AttendanceHandler) Clockout(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := ctx.Value(auth.CtxKeyAuthUserID).(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unauthorized user")))
		return
	}

	createdAttendance, err := h.attendanceUseCase.ClockOut(ctx, userID)
	if err != nil {
		switch true {
		case errors.Is(err, libs.ErrShouldClockIn{}):
			c.JSON(http.StatusUnprocessableEntity, httpresp.Err(err))
		default:
			c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		}

		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(dto.ClockinResponse{
		ID:         createdAttendance.ID,
		Date:       createdAttendance.Date.Format(time.DateOnly),
		ClockinAt:  createdAttendance.ClockinAt,
		ClockoutAt: createdAttendance.ClockoutAt,
	}))
}
