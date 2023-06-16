package server

import (
	"github.com/gin-gonic/gin"
	"github.com/znddzxx112/go-budd/service/graph_verification"
)

/**
* 图片验证码模块
**/
func (ds *defaultServer) graphGroupRouter() {
	graph := ds.engine.Group("/graph")
	{
		graph.GET("/code", ds.graphCode)
	}
}

//图片验证码
// @Description 图片验证码
// @Summary 图片验证码
// @Author znddzxx112@163.com
// @Date 22/12/20
// @Tags 图片验证码模块
// @Accept  image/jpeg
// @Router /graph/code [GET]
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) graphCode(ctx *gin.Context) {
	gv := graph_verification.NewGraphVerification(s.db)
	gv.ProduceGraphCodeAndSave(ctx.Writer)
}
