package exceptions

type InvalidDuration struct {
	ErrMessage string
}

func (e InvalidDuration) Error() string {
	return e.ErrMessage
}
func (e InvalidDuration) IsBusinessLogic() bool {
	return true
}
