package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"regexp"
)

// 生成密码
func Password(password string, salt string) string {
	sha := sha256.New()
	sha.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(sha.Sum([]byte(salt)))
}

// 检查密码格式
func CheckPassword(password string) bool {
	if ok, _ := regexp.MatchString("^[a-z0-9A-Z]{8,20}$", password); !ok {
		return false
	}
	if ok, _ := regexp.MatchString("[A-Z]{1,20}", password); !ok {
		return false
	}
	return true
}
