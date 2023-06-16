package server

import (
	"github.com/gin-gonic/gin"
	"github.com/znddzxx112/go-budd/pkg/utils"
	"github.com/znddzxx112/go-budd/service/user"
	"strconv"
)

/**
* 用户模块
**/
func (ds *defaultServer) userGroupRouter() {
	userNone := ds.engine.Group("/user")
	userNone.POST("/login", ds.UserLogin)
}

type UserLoginArgs struct {
	Mobile     string `json:"mobile" binding:"required" validate:"len=11" example:"18800011122"`
	Password   string `json:"password" example:"123456"`
	ImageToken string `json:"image_token"`
	ImageCode  string `json:"image_code"`
}
type UserLoginVO struct {
	Token string `json:"token" example:"TGT-FwLZ7dHoHepX4zMgVdB3WGmqZ94rMsK4"`
	St    string `json:"st" example:""`
}

type UserLoginVOResult struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Result  UserLoginVO `json:"result"`  // 数据结果集
}

// @Description 获取用户信息接口
// @Summary 获取用户信息接口
// @Author znddzxx112@163.com
// @Date 08/23/20
// @Tags 用户模块
// @Accept  json
// @Produce json
// @Router /user/login [post]
// @Param account body UserLoginArgs true "账号"
// @Success 200 {object} UserLoginVOResult "登录信息"
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) UserLogin(ctx *gin.Context) {
	args := new(UserLoginArgs)
	if err := ctx.ShouldBind(args); err != nil {
		// case101100
		utils.InvalidParametersError(ctx, err)
		return
	}
	if err := ValidatorInstance().Struct(args); err != nil {
		// case101101
		utils.InvalidParametersCheckError(ctx, err)
		return
	}

	vo := UserLoginVO{}
	userService := user.NewUserService(s.db)
	userId, err := userService.UserLogin(args.Mobile, args.Password, "WASDQER#@!")
	if err != nil {
		// case101102
		utils.ResponseError(ctx, 1, err.Error())
		return
	}

	// 创建session

	// case101103
	vo.Token = strconv.Itoa(userId)
	utils.ResponseSuccess(ctx, vo)
	return
}
