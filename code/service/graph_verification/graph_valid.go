package graph_verification

import (
	"github.com/jinzhu/gorm"
	"time"
)

/******sql******
CREATE TABLE `graph_valid` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `code` varchar(20) NOT NULL DEFAULT '' COMMENT '图形验证码',
  `token` varchar(100) NOT NULL DEFAULT '' COMMENT '凭证',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/

type GraphValid struct {
	ID        int       `gorm:"primary_key" json:"-"` // 主键
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	Code      string    `json:"code"`                 // 图形验证码
	Token     string    `json:"token"`                // 凭证
}

// TableName get sql table name.获取数据库表名
func (m *GraphValid) TableName() string {
	return "graph_valid"
}

func NewGraphValid() *GraphValid {
	return newGraphValid()
}

func NewGraphValidWithToken(db *gorm.DB, token string) *GraphValid {
	gv := newGraphValid()
	gv.Token = token
	notFound := db.Table(gv.TableName()).
		Select("*").
		Where("token = ?", token).
		First(gv).
		RecordNotFound()
	if notFound {
		return nil
	}
	return gv
}

func newGraphValid() *GraphValid {
	um := new(GraphValid)
	return um
}

//写入数据库
func (m *GraphValid) Create(db *gorm.DB) error {
	return db.Table(m.TableName()).Create(m).Error
}

func (m *GraphValid) Update(db *gorm.DB, update map[string]interface{}) error {
	return db.Table(m.TableName()).Where("id=?", m.ID).Update(update).Error
}

func (m *GraphValid) Delete(db *gorm.DB, delete map[string]interface{}) error {
	return db.Table(m.TableName()).Delete(GraphValid{}, delete).Error
}
