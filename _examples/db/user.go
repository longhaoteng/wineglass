package db

import (
	"gorm.io/gorm"

	"github.com/longhaoteng/wineglass/db"
)

type User struct {
	db.Model
	Name     string `json:"name" gorm:"type:string; size:32; index:idx_name; not null; comment:名称"`
	Password string `json:"-" gorm:"type:string; size:64; not null; comment:密码"`
	Gender   uint8  `json:"gender" gorm:"type:uint; comment:性别"`
	State    bool   `json:"state" gorm:"comment:状态"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	return nil
}

func (u *User) AfterCreate(db *gorm.DB) error {
	return nil
}

func (u *User) AfterSave(db *gorm.DB) error {
	return nil
}

func init() {
	db.AddRepositories(&User{})
}
