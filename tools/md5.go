package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// MD5 加密算法

// MD5 默认方法
func MD5(data string) string {
	return MD5String(data)
}

// MD5StringSimple 简洁的 md5 方式加密字符串
func MD5StringSimple(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// MD5String 完整的 md5 方式加密字符串
func MD5String(data string) string {
	has := md5.New()
	has.Write([]byte(data))
	b := has.Sum(nil)
	return hex.EncodeToString(b)
}

// MD5StringToUpper md5 加密后转换为大写
func MD5StringToUpper(data string) string {
	return strings.ToUpper(MD5String(data))
}

// MD5StringSimpleToUpper md5 加密后转换为大写
func MD5StringSimpleToUpper(data string) string {
	return strings.ToUpper(MD5StringSimple(data))
}

// MD5StringToLower md5 加密后转换为小写
func MD5StringToLower(data string) string {
	return strings.ToLower(MD5String(data))
}

// MD5StringSimpleToLower md5 加密后转换为小写
func MD5StringSimpleToLower(data string) string {
	return strings.ToLower(MD5StringSimple(data))
}
