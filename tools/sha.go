package tools

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// SHA256Simple 简单的sha256 加密
func SHA256Simple(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

// SHA256 SHA 256 加密
func SHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// SHA256String sha 256 加密字符串， 并返回字符串
func SHA256String(data string) string {
	return fmt.Sprintf("%x", SHA256([]byte(data)))
}

// DingDingSign timestamp + "\n" + appSecret当做签名字符串，使用HmacSHA256算法计算签名，然后进行Base64 encode，得到最终的签名值。
func DingDingSign(timestamp int64, appSecret string) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, appSecret)
	return base64.StdEncoding.EncodeToString(SHA256([]byte(stringToSign)))
}
