package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var exeDB *gorm.DB

// 打开配置文件
func openConfig() *ini.File {

	var err error
	work_path, _ := os.Getwd()
	path := strings.Split(work_path, "DangoMark")[0]
	config_path := filepath.Join(path, "/DangoMark/config", "/db.ini")
	config_file, err := ini.Load(config_path)
	if err != nil {
		log.Fatalf("读取配置文件 %s 失败: %s\n", config_path, err)
	}

	return config_file
}

// 连接数据库
func OpenDB() {

	config_file := openConfig()
	cfgSection := config_file.Section("DangoMark")
	url := cfgSection.Key("url").Value()
	username := cfgSection.Key("username").Value()
	password := cfgSection.Key("password").Value()
	connectInfo := fmt.Sprintf("%s:%s@%s", username, password, url)

	var err error
	exeDB, err = gorm.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("连接数据库 %s 失败: %s\n", "DangoMark", err)
	}
}

// 关闭数据库
func CloseDB() {
	if exeDB != nil {
		exeDB.Close()
	}
}
