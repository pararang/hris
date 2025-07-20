package libs

type ErrWeekendNotAllowed struct{}

func (e ErrWeekendNotAllowed) Error() string {
	return "clocking in on weekends is not allowed"
}
