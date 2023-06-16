package server

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/znddzxx112/go-budd/service/qqwry"
	"strings"
)

type defaultServerMock struct {
	defaultServer
	sqlMock sqlmock.Sqlmock
}

func NewDefaultServerMock(name string) Server {
	s := new(defaultServerMock)
	s.name = name
	return s
}

func (dsm *defaultServerMock) Run(port string, conf string) error {

	// 1、加载配置
	if err := dsm.config(conf); err != nil {
		return fmt.Errorf("rs.config(): %s", err.Error())
	}

	// 2、连接mysql
	if err := dsm.dbClient(); err != nil {
		return fmt.Errorf("rs.dbClient(): %s", err.Error())
	}

	// 3、mysql版本迁移
	//if err := tds.migrate(); err != nil {
	//	return fmt.Errorf("rs.migrate(): %s", err.Error())
	//}

	// 4、加载路由
	if err := dsm.router(); err != nil {
		return fmt.Errorf("rs.router(): %s", err.Error())
	}

	// 5、服务器初始化
	if err := dsm.init(); err != nil {
		return fmt.Errorf("rs.init(): %s", err.Error())
	}

	// 6、服务器监听端口
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return dsm.engine.Run(port)
}

func (dsm *defaultServerMock) Close() error {
	if dsm.db != nil {
		if err := dsm.db.Close(); err != nil {
			return err
		}
	}

	if qqwry.DefaultQQwry != nil {
		if err := qqwry.DefaultQQwry.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (dsm *defaultServerMock) dbClient() error {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return fmt.Errorf("dbClient(): sqlmock :%s", err.Error())
	}
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return fmt.Errorf("dbClient(): gorm :%s", err.Error())
	}
	dsm.db = gdb
	dsm.sqlMock = mock

	dsm.healthMock()
	dsm.userMock()
	return nil
}
