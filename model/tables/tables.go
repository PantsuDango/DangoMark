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
