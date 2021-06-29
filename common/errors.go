package common

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func ErrorRequest(err error, code int) error {
	return &RequestError{
		StatusCode: code,
		Err:        err,
	}
}
