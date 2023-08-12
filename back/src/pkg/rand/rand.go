package rand

import (
	"math/rand"
	"strings"
)

const UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LowerCase = "abcdefghijklmnopqrstuvwxyz"
const Nums = "0123456789"

/*
RandString会返回一个随机字符串
length: 字符串长度
source: 字符串内字符的来源
*/
func RandString(length int, source string) string {
	sourceLength := len(source)
	if sourceLength == 0 {
		panic("source string shouldn't be empty")
	}
	if length <= 0 {
		panic("length shouldn't little than 1")
	}
	sb := strings.Builder{}
	for ; length > 0; length-- {
		sb.WriteByte(source[rand.Intn(sourceLength)])
	}
	return sb.String()
}
