package libs

type ErrWeekendNotAllowed struct{}

func (e ErrWeekendNotAllowed) Error() string {
	return "clocking in on weekends is not allowed"
}

type ErrShouldClockIn struct{}

func (e ErrShouldClockIn) Error() string {
	return "you should clock in first"
}
