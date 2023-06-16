package utils

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// 设置日志格式
func SetLogLevel() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "level",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	logrus.SetLevel(logrus.InfoLevel)
}

// 设置日志时区
func SetDefaultLocation() {
	time.Local = time.FixedZone("UTC", 8*3600)
}

// 欢迎日志
func WelcomeLog(name, version string) {
	logrus.Printf("name: %s; version: %s, start...", name, version)
}

// 退出日志
func ExitLog(name, version string) {
	logrus.Printf("name: %s; version: %s, exit", name, version)
}

// 请求日志
func LogRequest(req *http.Request) {
	logrus.Info("path: ", req.URL.Path)
	logrus.Info(req.Header.Get("Authorization"))
	logrus.Info(req.Header.Get("cookie"))
	logrus.Info(req.Header.Get("Content-Type"))
}
