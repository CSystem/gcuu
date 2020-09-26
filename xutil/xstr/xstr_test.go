package xstr

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestRandStringBytes(t *testing.T) {
	for i := 1; i <= 10; i++ {
		slat := RandStringBytes(6)
		t.Log(slat)

		assert.Equal(t, len(slat), 6)
	}
}

func TestRandStringBytesMaskImprSrcUnsafe(t *testing.T) {
	for i := 1; i <= 10; i++ {
		slat := RandStringBytesMaskImprSrcUnsafe(6)
		t.Log(slat)

		assert.Equal(t, len(slat), 6)
	}
}

func TestMaskPhoneNumber(t *testing.T) {
	assert.Equal(t, MaskPhoneNumber("13200010002"), "132****0002")
	assert.Equal(t, MaskPhoneNumber("0"), "0")
	assert.Equal(t, MaskPhoneNumber("1234"), "1234")
}

func TestMaskRealname(t *testing.T) {
	assert.Equal(t, MaskRealname("老王", true), "老*")
	assert.Equal(t, MaskRealname("刘力气", true), "刘*")
	assert.Equal(t, MaskRealname("刘力气", false), "刘力气")
}
