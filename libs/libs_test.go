package libs

import (
	"testing"
	"time"
)

func TestIsWeekend(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Friday",
			args: args{date: time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "Saturday",
			args: args{date: time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "Sunday",
			args: args{date: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWeekend(tt.args.date); got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}
