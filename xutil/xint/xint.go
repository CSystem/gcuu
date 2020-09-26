package xint

import (
	"crypto/rand"
	"io"
	"strconv"
)

func RandInt(min, length int) int {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)

	if n != length {
		panic(err)
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	number, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	if number < min {
		return number + min
	}

	return number
}
