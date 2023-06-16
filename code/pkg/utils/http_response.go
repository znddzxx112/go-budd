package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ApiResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Result  interface{} `json:"result"`  // 数据结果集
}

// 操作成功
func ResponseSuccess(c *gin.Context, result interface{}) {
	responseOutput(c, 0, "操作成功", result)
}

// 错误
func ResponseError(c *gin.Context, code int, message string) {
	logrus.Errorf("错误code, %d", code)
	responseOutput(c, code, message, nil)
}

// 参数不合法
func InvalidParametersError(c *gin.Context, err error) {
	logrus.Errorf("参数不合法, %s", err.Error())
	responseError(c, 1, "参数不合法")
}

// 内部错误
func InternalServiceError(c *gin.Context, err error) {
	logrus.Errorf("内部错误, %s", err.Error())
	responseError(c, 2, "内部错误")
}

// 无权限访问
func NoPermissionError(c *gin.Context, err error) {
	logrus.Errorf("无权限访问, %s", err.Error())
	responseError(c, 3, "无权限访问")
}

// 参数校验不通过
func InvalidParametersCheckError(c *gin.Context, err error) {
	logrus.Errorf("参数校验不通过, %s", err.Error())
	responseError(c, 4, "参数校验不通过")
}

func responseOutput(c *gin.Context, code int, message string, result interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: message,
		Result:  result,
	})
}

func responseError(c *gin.Context, code int, message string) {
	responseOutput(c, code, message, nil)
}

type ApiRawResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Result  interface{} `json:"result"`  // 数据结果集
}

// raw输出
func ResponseRawError(w http.ResponseWriter, code int, message string, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	i := ApiRawResponse{
		Code:    code,
		Message: message,
		Result:  result,
	}
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	return
}
