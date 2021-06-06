package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

var (
	log_file    *os.File
	config_file *ini.File
	exeDB       *gorm.DB
	file_map    = make(map[string]map[string]string)
)

func init() {
	setLogFile()
	openConfig()
	openDB()
}

// 配置log日志
func setLogFile() {

	var err error
	work_path, _ := os.Getwd()
	path := strings.Split(work_path, "DangoMark")[0]
	log_path := filepath.Join(path, "/DangoMark/logs", "/update_file.log")
	log_file, err = os.OpenFile(log_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(log_file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// 打开配置文件
func openConfig() {

	var err error
	work_path, _ := os.Getwd()
	path := strings.Split(work_path, "DangoMark")[0]
	config_path := filepath.Join(path, "/DangoMark/config", "/db.ini")
	config_file, err = ini.Load(config_path)
	if err != nil {
		log.Fatalf("读取配置文件 %s 失败: %s\n", config_path, err)
	}
}

// 连接数据库
func openDB() {

	var err error
	cfgSection := config_file.Section("DangoMark")
	url := cfgSection.Key("url").Value()
	username := cfgSection.Key("username").Value()
	password := cfgSection.Key("password").Value()
	connectInfo := fmt.Sprintf("%s:%s@%s", username, password, url)
	exeDB, err = gorm.Open("mysql", connectInfo)

	if err != nil {
		log.Fatalf("连接数据库 %s_db 失败: %s\n", "DangoMark", err)
	}
}

// 读取图片文件
func readImageFile(path string) {

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("读取图片文件夹 %v 失败: %v\n", path, err.Error())
	}
	for _, file := range fs {
		if file.IsDir() {
			continue
		}
		for _, file_type := range []string{".png", ".jpg", ".txt"} {
			if strings.HasSuffix(file.Name(), file_type) {
				file_name := strings.Replace(file.Name(), file_type, "", -1)
				_, ok := file_map[file_name]
				if ok {
					if file_type == ".txt" {
						file_map[file_name]["txt"] = path + "/" + file.Name()
					} else {
						file_map[file_name]["image"] = path + "/" + file.Name()
					}
				} else {
					file_map[file_name] = make(map[string]string)
					if file_type == ".txt" {
						file_map[file_name]["txt"] = path + "/" + file.Name()
						file_map[file_name]["image"] = ""
					} else {
						file_map[file_name]["image"] = path + "/" + file.Name()
						file_map[file_name]["txt"] = ""
					}
				}
			}
		}
	} // 结束循环
}

// 图片转换为base64
func ImageToBase64(path string) string {
	image, _ := ioutil.ReadFile(path)
	imageBase64 := base64.StdEncoding.EncodeToString(image)
	return imageBase64
}

func main() {

	defer exeDB.Close()
	path := os.Args[1]
	readImageFile(path)
}
