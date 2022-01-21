package db

import "github.com/longhaoteng/wineglass/db"

func (u *User) Create() error {
	return db.DB.Create(&u).Error
}

func (u *User) IsExistsByName() (bool, error) {
	var count int64
	err := db.DB.Where("name = ?", u.Name).Count(&count).Error
	return count > 0, err
}
