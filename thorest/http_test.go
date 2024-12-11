package thorest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpError_Is(t *testing.T) {
	err := &HttpError{Code: 404}
	assert.True(t, errors.Is(err, ErrNotFound))
}
