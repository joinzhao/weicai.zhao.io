package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// AES 加密/解密

// AesEncrypt 默认的aes加密 默认使用 aes cbc
func AesEncrypt(data, key string) string {
	return fmt.Sprintf("%x", AesEncryptCBC([]byte(data), []byte(key)))
}

// AesDecrypt 默认的aes解密 默认使用 aes cbc
func AesDecrypt(data, key string) string {
	return fmt.Sprintf("%x", AesDecryptCBC([]byte(data), []byte(key)))
}

// AesEncryptCBC aes CBC 加密
func AesEncryptCBC(data, key []byte) []byte {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	data = pkcs5Padding(data, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 创建数组
	encrypted := make([]byte, len(data))
	// 加密
	blockMode.CryptBlocks(encrypted, data)
	return encrypted
}

// AesDecryptCBC aes CBC 解密
func AesDecryptCBC(encrypted, key []byte) []byte {
	// 分组秘钥
	block, _ := aes.NewCipher(key)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// 创建数组
	decrypted := make([]byte, len(encrypted))
	// 解密
	blockMode.CryptBlocks(decrypted, encrypted)
	// 去除补全码
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted
}

// pkcs5Padding 填充密钥
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// pkcs5UnPadding 去除填充
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// AesEncryptECB aes ECB 加密
func AesEncryptECB(data, key []byte) (encrypted []byte) {
	cip, _ := aes.NewCipher(generateKey(key))
	length := (len(data) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, data)
	pad := byte(len(plain) - len(data))
	for i := len(data); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cip.BlockSize(); bs <= len(data); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// AesDecryptECB aes ECB 解密
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cip, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))

	for bs, be := 0, cip.BlockSize(); bs < len(encrypted); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// generateKey 生成 ECB 模式下的key
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// AesEncryptCFB aes CFB 加密
func AesEncryptCFB(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)
	return encrypted, nil
}

// AesDecryptCFB aes CFB 解密
func AesDecryptCFB(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	var decrypted = make([]byte, len(data))
	stream.XORKeyStream(decrypted, data)
	return decrypted, nil
}
