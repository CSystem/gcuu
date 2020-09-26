package xstr

import (
	"math/rand"
	"regexp"
	"time"
	"unsafe"
)

// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandStringBytes - 生成指定长度的随机字符串
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
// https://colobu.com/2018/09/02/generate-random-string-in-Go/
func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano()) // 如果取消 seed 则生成的字符串顺序为固定值。

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

var src = rand.NewSource(time.Now().UnixNano())

// RandStringBytesMaskImprSrcUnsafe - 比 RandStringBytes 快 6.3 倍，并且只使用了六分之一的内存和一半的内存分配
// https://www.flysnow.org/2019/09/30/how-to-generate-a-random-string-of-a-fixed-length-in-go.html
func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// MaskPhoneNumber - 手机号掩码为 132****1111 的格式
func MaskPhoneNumber(number string) string {
	// 定义正则表达式
	r, _ := regexp.Compile("(^[0-9]{3})[0-9]{4}([0-9]{4})")

	// 替换后的结果
	result := r.ReplaceAllString(number, "$1****$2")

	return result
}

// MaskRealname - 汉字名称只保留姓
func MaskRealname(name string, mask bool) string {
	if !mask {
		return name
	}
	// 定义正则表达式
	r, _ := regexp.Compile(`([\p{Han}])([\p{Han}]+)`)

	// 替换后的结果
	result := r.ReplaceAllString(name, "$1*")

	return result
}
