package mapHttpErrors

type ErrorHTTPMapper interface {
	MapError(err error) (statusCode int, message string)
}
