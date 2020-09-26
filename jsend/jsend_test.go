package jsend

import (
	"net/http"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetStatus(t *testing.T) {
	code := getStatus(http.StatusOK)
	assert.Equal(t, StatusSuccess, code)

	code = getStatus(http.StatusBadRequest)
	assert.Equal(t, StatusFail, code)

	code = getStatus(http.StatusBadGateway)
	assert.Equal(t, StatusError, code)
}
