package usecase

import (
	"context"
	"fmt"

	"github.com/pararang/hris/domain/entity"
	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/libs/auth"
)

type reimbursementUseCase struct {
	reimbursementRepo repository.ReimbursementRepository
	attendanceRepo    repository.AttendanceRepository
}

func NewReimbursementUseCase(
	reimbursementRepo repository.ReimbursementRepository,
	attendanceRepo repository.AttendanceRepository,
) *reimbursementUseCase {
	return &reimbursementUseCase{
		reimbursementRepo: reimbursementRepo,
		attendanceRepo:    attendanceRepo,
	}
}

func (r *reimbursementUseCase) SubmitReimbursement(ctx context.Context, param dto.SubmitReimbursementParam) (*entity.Reimbursement, error) {
	createdBy, ok := ctx.Value(auth.CtxKeyAuthUserEmail).(string)
	if !ok {
		return nil, fmt.Errorf("error on get createdBy")
	}

	payrollPeriod, err := r.attendanceRepo.FindLatestPayrollPeriod(ctx)
	if err != nil {
		return nil, fmt.Errorf("error on find payroll period: %w", err)
	}

	data, err := r.reimbursementRepo.CreateReimbursement(ctx, &entity.Reimbursement{
		UserID:          param.UserID,
		Amount:          param.Amount,
		Description:     param.Description,
		TransactionDate: param.TransactionDate,
		PayrollPeriodID: payrollPeriod.ID,
		Status:          entity.StatusReimbursementPending,
		BaseModel: entity.BaseModel{
			CreatedBy: createdBy,
		},
	})

	return data, nil
}
