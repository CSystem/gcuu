package xstr

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestUUIDNumber(t *testing.T) {
	for i := 1; i <= 100; i++ {
		uuid := UUIDNumber()
		assert.Equal(t, len(uuid), 32)
	}
}
