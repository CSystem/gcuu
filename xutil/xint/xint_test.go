package xint

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

// 可以手动测试随机方法的数值，为了避免工具的缓存添加 -count=1 参数
// Read More: https://www.jianshu.com/p/dd69df14526d

// go test -count=1 -v github.com/wjp2013/LuckyCat/pkg/xutil/xint -test.run TestRandInt
func TestRandInt(t *testing.T) {
	for i := 1; i <= 100; i++ {
		randInt := RandInt(100000, 6)
		t.Log(randInt)
		assert.Equal(t, randInt > 100000, true)
	}
}

// go test -count=1 -v github.com/wjp2013/LuckyCat/pkg/xutil/xint -test.run TestCalculateRedPacketProbability
func TestCalculateRedPacketProbability(t *testing.T) {
	for i := 1; i <= 100; i++ {
		slat := CalculateRedPacketProbability(10)
		t.Log(slat)
	}
}

// go test -count=1 -v github.com/wjp2013/LuckyCat/pkg/xutil/xint -test.run TestCalculateRedPacketRatio
func TestCalculateRedPacketRatio(t *testing.T) {
	for i := 1; i <= 100; i++ {
		slat := CalculateRedPacketRatio(10)
		t.Log(slat)
	}
}
