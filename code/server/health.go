package server

import (
	"github.com/gin-gonic/gin"
	"github.com/znddzxx112/go-budd/pkg/utils"
)

/**
* 健康检查模块
**/
func (ds *defaultServer) healthGroupRouter() {
	pingNone := ds.engine.Group("/health")
	pingNone.POST("", ds.Pong)
}

type PongArgs struct {
	Say string `json:"say" binding:"required" validate:"gt=2"`
}

type PongVO struct {
	Res string `json:"res"`
}

// swagger::model PongVOResult
type PongVOResult struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 状态短语
	Result  PongVO `json:"result"`  // 数据结果集
}

// @Description 应用存活接口
// @Summary 应用存活接口
// @Author znddzxx112@163.com
// @Date 08/23/20
// @Tags 健康模块
// @Accept  json
// @Produce json
// @Router /health [post]
// @Param args body PongArgs true "入参"
// @Success 200 {object} PongVOResult "返回结果"
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) Pong(ctx *gin.Context) {
	args := new(PongArgs)
	if err := ctx.ShouldBindJSON(args); err != nil {
		// case100100
		utils.InvalidParametersError(ctx, err)
		return
	}
	if err := ValidatorInstance().Struct(args); err != nil {
		// case100101
		utils.InvalidParametersCheckError(ctx, err)
		return
	}
	pongVO := PongVO{
		Res: args.Say,
	}
	// case100102
	utils.ResponseSuccess(ctx, pongVO)
}
