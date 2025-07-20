package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/dto"
	"github.com/prrng/dealls/libs"
	"github.com/prrng/dealls/libs/auth"
	"github.com/prrng/dealls/libs/httpresp"
)

type OvertimeHandler struct {
	overtimeUseCase usecase.OvertimeUseCase
}

func NewOvertimeHandler(overtimeUseCase usecase.OvertimeUseCase) *OvertimeHandler {
	return &OvertimeHandler{
		overtimeUseCase: overtimeUseCase,
	}
}

func (h *OvertimeHandler) SubmitOvertime(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.SubmitOvertimeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	userID, ok := ctx.Value(auth.CtxKeyAuthUserID).(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unauthorized user")))
		return
	}

	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(errors.New("invalid date format")))
		return
	}

	overtime, err := h.overtimeUseCase.SubmitOvertime(ctx, &usecase.SubmitOvertimeParam{
		UserID:     userID,
		HoursTaken: req.HoursTaken,
		Date:       date,
		Reason:     req.Reason,
	})
	if err != nil {
		switch true {
		case errors.Is(err, libs.ErrShouldClockIn{}) || errors.Is(err, libs.ErrShouldClockOut{}):
			c.JSON(http.StatusUnprocessableEntity, httpresp.Err(fmt.Errorf("should attend for overtime: %w", err)))
		case errors.Is(err, libs.ErrOvertimeAlreadySubmitted{}):
			c.JSON(http.StatusConflict, httpresp.Err(err))
		default:
			c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		}

		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(dto.OvertimeResponse{
		ID:         overtime.ID,
		HoursTaken: overtime.HoursTaken,
		Date:       overtime.Date,
		Reason:     overtime.Reason,
		Status:     overtime.Status,
	}))
}
