package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/entity"
	"github.com/pararang/hris/libs"
	"github.com/pararang/hris/libs/auth"
)

type overtimeUseCase struct {
	overtimeRepo   repository.OvertimeRepository
	attendanceRepo repository.AttendanceRepository
}

// NewOvertimeUseCase creates a new instance of OvertimeUseCase
func NewOvertimeUseCase(
	overtimeRepo repository.OvertimeRepository,
	attendanceRepo repository.AttendanceRepository,
) *overtimeUseCase {
	return &overtimeUseCase{
		overtimeRepo:   overtimeRepo,
		attendanceRepo: attendanceRepo,
	}
}

func (o *overtimeUseCase) SubmitOvertime(ctx context.Context, param *dto.SubmitOvertimeParam) (*entity.Overtime, error) {
	_, err := o.overtimeRepo.FindUserOvertimeByDate(ctx, param.UserID, param.Date)
	if err == nil {
		return nil, libs.ErrOvertimeAlreadySubmitted{}
	}

	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("error on existing overtime: %w", err)
	}

	// only can submit ot:
	// -weekend ?
	// -weekday after work/clockout

	// if !libs.IsWeekend(param.Date) {
	attendanceToday, err := o.attendanceRepo.FindUserAttendanceByDate(ctx, param.UserID, param.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, libs.ErrShouldClockIn{}
		}

		return nil, fmt.Errorf("error on get today attendance: %w", err)
	}

	if attendanceToday.ClockoutAt == nil {
		return nil, libs.ErrShouldClockOut{}
	}
	// }

	createdBy, ok := ctx.Value(auth.CtxKeyAuthUserEmail).(string)
	if !ok {
		return nil, fmt.Errorf("error on get createdBy")
	}

	payrollPeriod, err := o.attendanceRepo.FindLatestPayrollPeriod(ctx)
	if err != nil {
		return nil, fmt.Errorf("error on find payroll period: %w", err)
	}

	overtime, err := o.overtimeRepo.CreateOvertime(ctx, &entity.Overtime{
		UserID:          param.UserID,
		Date:            param.Date,
		HoursTaken:      param.HoursTaken,
		Reason:          param.Reason,
		Status:          entity.StatusOvertimePending,
		PayrollPeriodID: payrollPeriod.ID,
		BaseModel: entity.BaseModel{
			CreatedBy: createdBy,
			UpdatedAt: time.Now(),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error on create overtime: %w", err)
	}

	return overtime, nil
}
