package xstr

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// 通过 UUID v4 的方案生成微信商户订单号
// 商户订单号，需保持唯一性，只能是字母或者数字，不能包含有其他字符，长度为32位。
func UUIDNumber() (str string) {
	str = uuid.NewV4().String()
	return strings.Replace(str, "-", "", -1)
}
