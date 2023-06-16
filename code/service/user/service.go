package user

import (
	"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	// 用户模块
	//
	// 用户登录
	UserLogin(mobile, password, SecretKey string) (int, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	u := new(userService)
	u.db = db
	return u
}

func (us *userService) UserLogin(mobile, password, SecretKey string) (int, error) {
	user := NewUserModelWithMobile(us.db, mobile)
	if user == nil {
		// case101102
		return 0, fmt.Errorf("%s", "用戶不存在或密码错误")
	}

	// check password

	// case101103
	return user.ID, nil
}
