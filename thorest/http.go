package thorest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"states"`
	Path    string `json:"path"`
}

var ErrNotFound = &HttpError{Code: 404, Status: "not found", Message: "resource not found"}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) String() string {
	if e.Path != "" {
		return fmt.Sprintf("HTTP(%s) %d: %s", e.Path, e.Code, e.Message)
	}
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
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
