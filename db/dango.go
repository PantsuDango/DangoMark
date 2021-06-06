package db

import (
	"DangoMark/model/tables"
)

type SocialDB struct{}

func (SocialDB) GetUserInfo(User string) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("user = ?", User).First(&user).Error
	return user, err
}

func (SocialDB) QueryUserById(Id int) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("id = ?", Id).First(&user).Error
	return user, err
}
