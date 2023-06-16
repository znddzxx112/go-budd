package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserModel struct {
	ID           int       `gorm:"primary_key" json:"-"`
	CreatedAt    time.Time `json:"created_at"`     // 创建时间
	Account      string    `json:"account"`        // 账户名称
	TrueName     string    `json:"true_name"`      // 真实姓名
	Mobile       string    `json:"mobile"`         // 手机号码
	Password     string    `json:"password"`       // 登录密码
	Status       int8      `json:"status"`         // 1:可用 2：已删除
	LastLogin    time.Time `json:"last_login"`     // 上次登录时间
	ThisLogin    time.Time `json:"this_login"`     // 本次登录时间
	CreateUserId int       `json:"create_user_id"` // 创建本用户的用户id
}

func NewUserModel() *UserModel {
	return newUserModel()
}

func (u *UserModel) TableName() string {
	return "user"
}

func NewUserModelWithId(db *gorm.DB, id int) *UserModel {
	r := newUserModel()
	r.ID = id
	notFound := db.Table(r.TableName()).
		Select("*").
		Where("id = ?", id).
		Where("status = 1").
		First(r).
		RecordNotFound()
	if notFound {
		return nil
	}
	return r
}

func NewUserModelWithMobile(db *gorm.DB, mobile string) *UserModel {
	r := newUserModel()
	r.Mobile = mobile
	notFound := db.Table(r.TableName()).
		Select("*").
		Where("mobile = ?", mobile).
		Where("status = ?", 1).
		First(r).
		RecordNotFound()
	if notFound {
		return nil
	}
	return r
}

func NewUserModelWithAccount(db *gorm.DB, account string) *UserModel {
	r := newUserModel()
	r.Account = account
	notFound := db.Table(r.TableName()).
		Select("*").
		Where("account = ?", account).
		Where("status = ?", 1).
		First(r).
		RecordNotFound()
	if notFound {
		return nil
	}
	return r
}

func newUserModel() *UserModel {
	um := new(UserModel)
	return um
}

func (u *UserModel) Create(db *gorm.DB) error {
	return db.Table(u.TableName()).Create(u).Error
}

func (u *UserModel) Update(db *gorm.DB, update map[string]interface{}) error {
	return db.Table(u.TableName()).Where("id=?", u.ID).Update(update).Error
}

func CountCreateUser(db *gorm.DB, userId int) int {
	userModel := newUserModel()
	var count int = 0
	err := db.Table(userModel.TableName()).Where("create_user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

type UserModelList []UserModel

func ListUserModelWithUserIds(db *gorm.DB, userIds []int, search string, limit, offset int) (*UserModelList, int, error) {
	listCap := limit
	if limit == 0 {
		listCap = 100
	}
	userModel := newUserModel()
	userModelList := make(UserModelList, 0, listCap)
	sql := db.Table(userModel.TableName()).
		Select("*").Where("status = ?", 1)
	if userIds != nil && len(userIds) > 0 {
		sql = sql.Where("id in (?)", userIds)
	}
	if search != "" {
		sql = sql.Where("true_name LIKE ? OR mobile LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	var total int
	if limit != 0 {
		sql = sql.Limit(limit)
	}
	err := sql.Count(&total).
		Offset(offset).
		Scan(&userModelList).Error
	return &userModelList, total, err
}

func ListUserModelWithAccounts(db *gorm.DB, accounts []string, limit, offset int) (UserModelList, int, error) {
	listCap := limit
	if limit == 0 {
		listCap = 100
	}
	userModel := newUserModel()
	userModelList := make(UserModelList, 0, listCap)
	sql := db.Table(userModel.TableName()).
		Select("*").Where("status = ?", 1)
	if accounts != nil && len(accounts) > 0 {
		sql = sql.Where("account in (?)", accounts)
	}
	var total int
	if limit != 0 {
		sql = sql.Limit(limit)
	}
	err := sql.Count(&total).
		Offset(offset).
		Scan(&userModelList).Error
	return userModelList, total, err
}

func (uList *UserModelList) GetUserModel(userId int) *UserModel {
	for _, v := range *uList {
		if v.ID == userId {
			return &v
		}
	}
	return nil
}

func (uList *UserModelList) GetUserModelWithAccount(account string) *UserModel {
	for _, v := range *uList {
		if v.Account == account {
			return &v
		}
	}
	return nil
}

func (uList *UserModelList) UserIds() []int {
	userIds := make([]int, 0, len(*uList))
	for _, v := range *uList {
		userIds = append(userIds, v.ID)
	}
	return userIds
}
