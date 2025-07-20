package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/dto"
	"github.com/prrng/dealls/libs/auth"
	"github.com/prrng/dealls/libs/httpresp"
)

type ReimbursementHandler struct {
	reimbursementUseCase usecase.ReimbursementUseCase
}

func NewReimbursementHandler(reimbursementUseCase usecase.ReimbursementUseCase) *ReimbursementHandler {
	return &ReimbursementHandler{
		reimbursementUseCase: reimbursementUseCase,
	}
}

func (h *ReimbursementHandler) SubmitReimbursement(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.SubmitReimbursementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	userID, ok := ctx.Value(auth.CtxKeyAuthUserID).(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unauthorized user")))
		return
	}

	date, err := time.Parse(time.DateOnly, req.TransactionDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(errors.New("invalid date format")))
		return
	}

	reimbursement, err := h.reimbursementUseCase.SubmitReimbursement(ctx, usecase.SubmitReimbursementParam{
		UserID:          userID,
		Amount:          float64(req.Amount),
		Description:     req.Description,
		TransactionDate: date,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(dto.ReimbursementResponse{
		ID:              reimbursement.ID,
		Amount:          int(reimbursement.Amount),
		Description:     reimbursement.Description,
		TransactionDate: reimbursement.TransactionDate.Format(time.DateOnly),
		Status:          reimbursement.Status,
	}))
}
