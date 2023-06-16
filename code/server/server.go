package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	dbStub "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/znddzxx112/go-budd/docs"
	"github.com/znddzxx112/go-budd/pkg/utils"
	"github.com/znddzxx112/go-budd/service/qqwry"
	"strings"
	"time"
)

type Server interface {
	Run(port string, conf string) error
	Close() error
}

type serverConfig struct {
	Debug         bool            `yaml:"debug"`
	DbConfig      *utils.DBConfig `required:"true" yaml:"mysql"`
	QQwryPath     string          `required:"true" yaml:"qqwry_path"`
	RsaPrivateKey string          `required:"true" yaml:"rsa_private_key"`
}

type defaultServer struct {
	name string

	conf   *serverConfig
	db     *gorm.DB
	engine *gin.Engine
}

func NewServer(name string) Server {
	s := new(defaultServer)
	s.name = name
	return s
}

// 服务器启动
func (ds *defaultServer) Run(port string, conf string) error {
	// 1、加载配置
	if err := ds.config(conf); err != nil {
		return fmt.Errorf("rs.config(): %s", err.Error())
	}

	// 2、连接mysql
	if err := ds.dbClient(); err != nil {
		return fmt.Errorf("rs.dbClient(): %s", err.Error())
	}

	// 3、mysql版本迁移
	if err := ds.migrate(); err != nil {
		return fmt.Errorf("rs.migrate(): %s", err.Error())
	}

	// 4、加载路由
	if err := ds.router(); err != nil {
		return fmt.Errorf("rs.router(): %s", err.Error())
	}

	// 5、服务器初始化
	if err := ds.init(); err != nil {
		return fmt.Errorf("rs.init(): %s", err.Error())
	}

	// 6、服务器监听端口
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return ds.engine.Run(port)
}

// 服务器关闭
func (ds *defaultServer) Close() error {

	if ds.db != nil {
		if err := ds.db.Close(); err != nil {
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

func (ds *defaultServer) config(configPath string) error {
	ds.conf = new(serverConfig)
	err := configor.Load(ds.conf, configPath)
	if err != nil {
		return err
	}
	return nil
}

func (ds *defaultServer) dbClient() error {
	db, err := utils.GetDBConnection(ds.conf.DbConfig)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}
	ds.db = db
	return nil
}

func (ds *defaultServer) migrate() error {
	db, err := utils.GetMigrationDBConnection(ds.conf.DbConfig)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}
	db.DB().SetMaxIdleConns(1)
	db.DB().SetMaxOpenConns(1)
	instance, err := dbStub.WithInstance(db.DB(), &dbStub.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://resources/migration", "mysql", instance)
	if err != nil {
		return err
	}
	beforeVersion, _, _ := m.Version()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	if afterVersion, dirty, err := m.Version(); err != nil {
		return err
	} else {
		logrus.Infof("before migrate: %+v, after migrate version: %+v, %+v", beforeVersion, afterVersion, dirty)
	}
	return nil
}

func (ds *defaultServer) router() error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	ds.engine = r
	if ds.conf.Debug {
		// swagger文档
		ds.swaggerDoc()
	}
	ds.graphGroupRouter()
	ds.healthGroupRouter()
	ds.userGroupRouter()
	return nil
}

// http://127.0.0.1:7890/swagger/index.html
func (ds *defaultServer) swaggerDoc() {
	docs.SwaggerInfo.Title = ds.name
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	ds.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (ds *defaultServer) init() error {
	var err error
	// 加载ip库
	qqwry.DefaultQQwry, err = qqwry.NewQQwry(ds.conf.QQwryPath)
	if err != nil {
		return fmt.Errorf("打开ip库:%s", err.Error())
	}

	// 每隔10s
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				logrus.Info("exec defaultServer.RefreshTrie()")

			}
		}
	}()
	return nil
}
