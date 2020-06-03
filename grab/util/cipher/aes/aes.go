package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	cipher2 "github.com/R0L/golang-test/grab/util/cipher"
)

type aesEntity struct {
	key []byte
}

// 构造函数
func NewAES(key string) cipher2.Cipher {
	return &aesEntity{
		key: []byte(key),
	}
}

// 加密
func (s *aesEntity) Encrypt(src string) (string, error) {
	byteSrc := []byte(src)
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}
	byteSrc = padding(byteSrc, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, s.key)
	blockMode.CryptBlocks(byteSrc, byteSrc)
	return base64.StdEncoding.EncodeToString(byteSrc), nil
}

// 解密
func (s *aesEntity) Decrypt(src string) (string, error) {
	byteSrc, err := base64.StdEncoding.DecodeString(src)
	if nil != err {
		return "", err
	}
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, s.key)
	blockMode.CryptBlocks(byteSrc, byteSrc)
	byteSrc = unpadding(byteSrc)
	return string(byteSrc[:]), nil
}

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := make([]byte, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	index := bytes.IndexByte(src, 0)
	return src[:index]
}
