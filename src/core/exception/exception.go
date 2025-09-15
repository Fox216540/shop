package exception

type Exception struct {
	Message string
}

func (e Exception) Error() string {
	return e.Message
}
