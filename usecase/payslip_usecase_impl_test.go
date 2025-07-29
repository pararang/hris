package usecase

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/entity"
	"github.com/pararang/hris/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_payslipUseCase_calculatePeriodeWorkingDay(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "weekdays only",
			args: args{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC),
			},
			want: 5,
		},
		{
			name: "including weekend",
			args: args{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
			},
			want: 5,
		},
		{
			name: "full month",
			args: args{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				endDate:   time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: 22, // 31 days minus weekends (9 days)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &payslipUseCase{}
			if got := p.calculatePeriodeWorkingDay(tt.args.startDate, tt.args.endDate); got != tt.want {
				t.Errorf("payslipUseCase.calculatePeriodeWorkingDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_payslipUseCase_calculateOvertimePay(t *testing.T) {
	type args struct {
		baseSalary    float64
		overtimeHours int32
		workingDays   int16
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic overtime calculation",
			args: args{
				baseSalary:    5000000,
				overtimeHours: 10,
				workingDays:   20,
			},
			want: 625000,
		},
		{
			name: "zero overtime hours",
			args: args{
				baseSalary:    5000000,
				overtimeHours: 0,
				workingDays:   20,
			},
			want: 0,
		},
		{
			name: "different working days",
			args: args{
				baseSalary:    5000000,
				overtimeHours: 10,
				workingDays:   20,
			},
			want: 625000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &payslipUseCase{}
			if got := p.calculateOvertimePay(tt.args.baseSalary, tt.args.overtimeHours, tt.args.workingDays); got != tt.want {
				t.Errorf("payslipUseCase.calculateOvertimePay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_payslipUseCase_calculateProrateBaseSalary(t *testing.T) {
	type args struct {
		baseSalary  float64
		attendDays  int16
		workingDays int16
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "full attendance",
			args: args{
				baseSalary:  5000000,
				attendDays:  20,
				workingDays: 20,
			},
			want: 5000000,
		},
		{
			name: "partial attendance",
			args: args{
				baseSalary:  5000000,
				attendDays:  10,
				workingDays: 20,
			},
			want: 2500000,
		},
		{
			name: "no attendance",
			args: args{
				baseSalary:  5000000,
				attendDays:  0,
				workingDays: 20,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &payslipUseCase{}
			if got := p.calculateProrateBaseSalary(tt.args.baseSalary, tt.args.attendDays, tt.args.workingDays); got != tt.want {
				t.Errorf("payslipUseCase.calculateProrateBaseSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_payslipUseCase_calculateTHPComponents(t *testing.T) {
	type fields struct {
		db                *sql.DB
		payslipRepo       repository.PayslipRepository
		userRepo          repository.UserRepository
		attendanceRepo    repository.AttendanceRepository
		overtimeRepo      repository.OvertimeRepository
		reimbursementRepo repository.ReimbursementRepository
	}
	type args struct {
		ctx   context.Context
		param calculateTHPParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.PayslipDetails
		wantErr bool
		setup   func(attendanceRepo *mocks.AttendanceRepository, overtimeRepo *mocks.OvertimeRepository, reimbursementRepo *mocks.ReimbursementRepository)
	}{
		{
			name:   "success with all components",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				param: calculateTHPParam{
					UserID:          uuid.New(),
					PayrollPeriodID: uuid.New(),
					BaseSalary:      5000000,
					WorkingDays:     20,
				},
			},
			want: entity.PayslipDetails{
				AttendanceDays:    15,
				ProrateBaseSalary: 3750000,
				OvertimePay:       625000,
				OvertimeHours:     10,
				ReimbursementPay:  200000,
			},
			wantErr: false,
			setup: func(attendanceRepo *mocks.AttendanceRepository, overtimeRepo *mocks.OvertimeRepository, reimbursementRepo *mocks.ReimbursementRepository) {
				attendanceRepo.On("CountAttendance", mock.Anything, mock.Anything, mock.Anything).Return(15, nil)
				overtimeRepo.On("CountUserOvertimeHoursInPeriod", mock.Anything, mock.Anything, mock.Anything).Return(int32(10), nil)
				reimbursementRepo.On("CountUserApprovedAmountReimbursementByPeriod", mock.Anything, mock.Anything, mock.Anything).Return(float64(200000), nil)
			},
		},
		{
			name:   "success with no attendance",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				param: calculateTHPParam{
					UserID:          uuid.New(),
					PayrollPeriodID: uuid.New(),
					BaseSalary:      5000000,
					WorkingDays:     20,
				},
			},
			want: entity.PayslipDetails{
				AttendanceDays:    0,
				ProrateBaseSalary: 0,
				OvertimePay:       625000,
				OvertimeHours:     10,
				ReimbursementPay:  200000,
			},
			wantErr: false,
			setup: func(attendanceRepo *mocks.AttendanceRepository, overtimeRepo *mocks.OvertimeRepository, reimbursementRepo *mocks.ReimbursementRepository) {
				attendanceRepo.On("CountAttendance", mock.Anything, mock.Anything, mock.Anything).Return(0, sql.ErrNoRows)
				overtimeRepo.On("CountUserOvertimeHoursInPeriod", mock.Anything, mock.Anything, mock.Anything).Return(int32(10), nil)
				reimbursementRepo.On("CountUserApprovedAmountReimbursementByPeriod", mock.Anything, mock.Anything, mock.Anything).Return(float64(200000), nil)
			},
		},
		{
			name:   "error on attendance",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				param: calculateTHPParam{
					UserID:          uuid.New(),
					PayrollPeriodID: uuid.New(),
					BaseSalary:      5000000,
					WorkingDays:     20,
				},
			},
			want:    entity.PayslipDetails{},
			wantErr: true,
			setup: func(attendanceRepo *mocks.AttendanceRepository, overtimeRepo *mocks.OvertimeRepository, reimbursementRepo *mocks.ReimbursementRepository) {
				attendanceRepo.On("CountAttendance", mock.Anything, mock.Anything, mock.Anything).Return(0, errors.New("database error"))
				overtimeRepo.On("CountUserOvertimeHoursInPeriod", mock.Anything, mock.Anything, mock.Anything).Return(int32(0), nil)
				reimbursementRepo.On("CountUserApprovedAmountReimbursementByPeriod", mock.Anything, mock.Anything, mock.Anything).Return(float64(0), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attendanceRepo := &mocks.AttendanceRepository{}
			overtimeRepo := &mocks.OvertimeRepository{}
			reimbursementRepo := &mocks.ReimbursementRepository{}

			if tt.setup != nil {
				tt.setup(attendanceRepo, overtimeRepo, reimbursementRepo)
			}

			p := &payslipUseCase{
				db:                tt.fields.db,
				payslipRepo:       tt.fields.payslipRepo,
				userRepo:          tt.fields.userRepo,
				attendanceRepo:    attendanceRepo,
				overtimeRepo:      overtimeRepo,
				reimbursementRepo: reimbursementRepo,
			}
			got, err := p.calculateTHPComponents(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("payslipUseCase.calculateTHPComponents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("payslipUseCase.calculateTHPComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
