package db

import (
	"DangoMark/model/tables"
)

type DangoDB struct{}

func (DangoDB) GetUserInfo(User string) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("user = ?", User).First(&user).Error
	return user, err
}

func (DangoDB) QueryUserById(Id int) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("id = ?", Id).First(&user).Error
	return user, err
}

func (DangoDB) SaveImageData(image_data tables.ImageData) {
	exeDB.Save(&image_data)
}
