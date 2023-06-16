package utils

import "regexp"

// 检查手机号码格式
func CheckPhone(phone string) bool {
	if ok, _ := regexp.MatchString("^[1]([3-9])[0-9]{9}$", phone); !ok {
		return false
	}
	return true
}
