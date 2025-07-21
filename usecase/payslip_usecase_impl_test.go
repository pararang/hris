package usecase

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/entity"
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &payslipUseCase{
				db:                tt.fields.db,
				payslipRepo:       tt.fields.payslipRepo,
				userRepo:          tt.fields.userRepo,
				attendanceRepo:    tt.fields.attendanceRepo,
				overtimeRepo:      tt.fields.overtimeRepo,
				reimbursementRepo: tt.fields.reimbursementRepo,
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
