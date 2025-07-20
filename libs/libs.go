package libs

import "time"

func IsWeekend(date time.Time) bool {
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
}
