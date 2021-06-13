package tables

import "time"

type User struct {
	ID        int       `json:"ID"          gorm:"column:id"`
	User      string    `json:"User"        gorm:"column:user"`
	Password  string    `json:"Password"    gorm:"column:password"`
	Ip        string    `json:"Ip"          gorm:"column:ip"`
	Total     string    `json:"Total"       gorm:"column:total"`
	CreatedAt time.Time `json:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  gorm:"column:lastupdate"`
}

func (User) TableName() string {
	return "user"
}

type ImageData struct {
	ID             int       `json:"ID"              gorm:"column:id"`
	Url            string    `json:"Url"             gorm:"column:url"`
	Language       string    `json:"Language"        gorm:"column:language"`
	Suggestion     string    `json:"Suggestion"      gorm:"column:suggestion"`
	MarkResult     string    `json:"MarkResult"      gorm:"column:mark_result"`
	QualityResult  string    `json:"QualityResult"   gorm:"column:quality_result"`
	CoordinateJson string    `json:"CoordinateJson"  gorm:"column:coordinate_json"`
	Status         int       `json:"Status"          gorm:"column:status"`
	CreatedAt      time.Time `json:"CreateTime"      gorm:"column:createtime"`
	UpdatedAt      time.Time `json:"UpdateTime"      gorm:"column:lastupdate"`
}

func (ImageData) TableName() string {
	return "image_data"
}
