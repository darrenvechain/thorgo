package thorest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string
	Status  string
	Path    string
}

func baseErr(code int) *HttpError {
	return &HttpError{
		Code:   code,
		Status: http.StatusText(code),
	}
}

var (
	ErrBadRequest         = baseErr(http.StatusBadRequest)
	ErrUnauthorized       = baseErr(http.StatusUnauthorized)
	ErrForbidden          = baseErr(http.StatusForbidden)
	ErrNotFound           = baseErr(http.StatusNotFound)
	ErrInternal           = baseErr(http.StatusInternalServerError)
	ErrServiceUnavailable = baseErr(http.StatusServiceUnavailable)
)

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) String() string {
	if e.Path != "" {
		return fmt.Sprintf("HTTP(%s) %s: %s", e.Path, e.Status, e.Message)
	}
	return fmt.Sprintf("HTTP %s: %s", e.Status, e.Message)
}

func newHttpError(resp *http.Response) *HttpError {
	message := resp.Status

	if resp.Body != nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err == nil {
			message = string(body)
		}
	}

	return &HttpError{
		Code:    resp.StatusCode,
		Status:  resp.Status,
		Message: message,
		Path:    resp.Request.URL.Path,
	}
}

func (e *HttpError) Is(target error) bool {
	var t *HttpError
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Code == t.Code
}
