package server

import "github.com/gin-gonic/gin"

/**
* ip地址模块
**/
func (ds *defaultServer) ipGroupRouter() {
	pingNone := ds.engine.Group("/ip")
	pingNone.POST("", ds.Pong)
}

func (s *defaultServer) Ip(ctx *gin.Context) {

}
