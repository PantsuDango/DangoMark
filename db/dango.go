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

func (DangoDB) SelectImageDataByLanguageAndStatus(language string, status int) (int, tables.ImageData) {
	var image_data tables.ImageData
	var count int
	exeDB.Model(&tables.ImageData{}).Where(`language = ? and status = ?`, language, status).Count(&count)
	exeDB.Where(`language = ? and status = ?`, language, status).Order("lastupdate ASC").Limit(1).Find(&image_data)
	return count, image_data
}
