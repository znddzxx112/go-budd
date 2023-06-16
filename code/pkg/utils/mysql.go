package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/url"
)

type DBConfig struct {
	User     string `default:"root" yaml:"user"`
	Password string `default:"" yaml:"password"`
	Name     string `yaml:"ip"`
	Port     uint   `default:"3306" yaml:"port"`
	DbName   string `required:"true" yaml:"db_name"`
	Charset  string `default:"utf8" yaml:"charset"`
	MaxIdle  int    `default:"10" yaml:"max_idle"`
	MaxOpen  int    `default:"50" yaml:"max_open"`
	LogMode  bool   `yaml:"log_mode"`
	Loc      string `required:"true" yaml:"loc"`
}

// 建立mysql连接池
func GetDBConnection(conf *DBConfig) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s"
	dsn := fmt.Sprintf(format, conf.User, conf.Password, conf.Name, conf.Port, conf.DbName, conf.Charset, url.QueryEscape(conf.Loc))
	logrus.Infof("dsn=%s", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(conf.LogMode)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	db.DB().SetMaxOpenConns(conf.MaxOpen)
	return db, nil
}

// 建立mysql连接池
func GetMigrationDBConnection(conf *DBConfig) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&multiStatements=true"
	dsn := fmt.Sprintf(format, conf.User, conf.Password, conf.Name, conf.Port, conf.DbName, conf.Charset, url.QueryEscape(conf.Loc))
	logrus.Infof("dsn=%s", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(conf.LogMode)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	db.DB().SetMaxOpenConns(conf.MaxOpen)
	return db, nil
}
