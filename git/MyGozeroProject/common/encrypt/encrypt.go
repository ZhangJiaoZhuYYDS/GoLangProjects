// @Author zhangjiaozhu 2023/9/28 23:34:00
package encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/zeromicro/go-zero/core/codec"
)

const (
	passwordEncryptSeed = "(beyond)@#$"
	mobileAesKey        = "5A2E746B08D846502F37A6E2D85D583B"
)

// 密码加密
func EncPassword(password string) string {
	return Md5Sum([]byte(strings.TrimSpace(password + passwordEncryptSeed)))
}

// 手机号加密
func EncMobile(mobile string) (string, error) {
	data, err := codec.EcbEncrypt([]byte(mobileAesKey), []byte(mobile))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// 手机号解密
func DecMobile(mobile string) (string, error) {
	originalData, err := base64.StdEncoding.DecodeString(mobile)
	if err != nil {
		return "", err
	}
	data, err := codec.EcbDecrypt([]byte(mobileAesKey), originalData)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Md5Sum(data []byte) string {
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}

func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp[16:]
}
