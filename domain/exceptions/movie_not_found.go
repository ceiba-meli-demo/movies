package exceptions

type MovieNotFound struct {
	ErrMessage string
}

func (e MovieNotFound) Error() string {
	return e.ErrMessage
}
