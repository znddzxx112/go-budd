package utils

import "time"

const TokenExpire = time.Hour * 12

// time转化为格式化时间
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 格式化时间转化为time
func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
}

// 时间戳转化为格式化时间
func FormatTimeWithInt(sec int) string {
	return time.Unix(int64(sec), 0).Format("2006-01-02 15:04:05")
}

// 时间戳转化为time
func ParseTimeWithInt(sec int) time.Time {
	return time.Unix(int64(sec), 0)
}
