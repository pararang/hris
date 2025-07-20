package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/dto"
	"github.com/prrng/dealls/libs/httpresp"
)

// PayslipHandler handles payslip-related HTTP requests
type PayslipHandler struct {
	payslipUseCase usecase.PayslipUseCase
}

// NewPayslipHandler creates a new instance of PayslipHandler
func NewPayslipHandler(payslipUseCase usecase.PayslipUseCase) *PayslipHandler {
	return &PayslipHandler{
		payslipUseCase: payslipUseCase,
	}
}

func (h *PayslipHandler) ProcessPayroll(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.ProcessPayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	err := h.payslipUseCase.GeneratePayslip(ctx, req.PayrollPeriodID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusCreated, httpresp.OK(nil))
}
