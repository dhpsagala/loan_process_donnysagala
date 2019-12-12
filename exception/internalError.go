package exception

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}
