package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pararang/hris/domain/usecase"
	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/entity"
	"github.com/pararang/hris/libs/auth"
	"github.com/pararang/hris/libs/httpresp"
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

	c.JSON(http.StatusCreated, httpresp.OK(gin.H{"message": "Payslip generated successfully"}))
}

func (h *PayslipHandler) ListUserPayslips(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := ctx.Value(auth.CtxKeyAuthUserID).(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("unauthorized user")))
		return
	}

	payslips, err := h.payslipUseCase.GetListPayslip(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.OK(h.buildPayslipListResponse(payslips)))
}

func (h *PayslipHandler) buildPayslipListResponse(payslips []entity.Payslip) (resp []dto.ItemListPayslipResponse) {
	for _, payslip := range payslips {
		resp = append(resp, dto.ItemListPayslipResponse{
			ID:            payslip.ID,
			GeneratedDate: payslip.CreatedAt.Format(time.DateTime),
			BaseSalary:    payslip.BaseSalary,
			Details:       payslip.DetailsJSON,
		})
	}
	return
}

func (h *PayslipHandler) GetPayslipDetails(c *gin.Context) {
	ctx := c.Request.Context()
	payslipID := uuid.MustParse(c.Param("id"))

	payslip, err := h.payslipUseCase.GetPayslipDetail(ctx, payslipID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.OK(payslip))
}

func (h *PayslipHandler) GetPayrollSummary(c *gin.Context) {
	ctx := c.Request.Context()
	periodID := uuid.MustParse(c.Param("id"))

	data, err := h.payslipUseCase.GetPayrollPeriodSummary(ctx, periodID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresp.Err(err))
		return
	}

	c.JSON(http.StatusOK, httpresp.OK(data))
}
