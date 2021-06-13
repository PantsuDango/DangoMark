package main

import (
	"DangoMark/model/params"
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/ffmt.v1"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

var (
	log_file   *os.File
	db_config  *ini.File
	pub_config *ini.File
	exeDB      *gorm.DB
	file_map   = make(map[string]map[string]string)
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

	work_path, _ := os.Getwd()
	path := strings.Split(work_path, "DangoMark")[0]
	db_config_path := filepath.Join(path, "/DangoMark/config", "/db.ini")
	db_config, _ = ini.Load(db_config_path)
	pub_config_path := filepath.Join(path, "/DangoMark/config", "/pub.ini")
	pub_config, _ = ini.Load(pub_config_path)

}

// 连接数据库
func openDB() {

	var err error
	cfgSection := db_config.Section("DangoMark")
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

func openFile() {

	dfsURL := pub_config.Section("filesystem").Key("fileserverupload").Value()
	for key, value := range file_map {
		for fileType, filePath := range value {
			fileContent, _ := ioutil.ReadFile(filePath) // just pass the file name
			switch fileType {
			case "image":
				fileURL := saveFileToDFS(dfsURL, fileContent)
				file_map[key]["image"] = fileURL
			case "txt":
				file_map[key]["text"] = string(fileContent)
			}
		}

	} // 结束循环
}

func saveFileToDFS(destURL string, fileContent []byte) string {

	// format params
	fields := map[string]string{
		"file":   uuid.NewV4().String(),
		"output": "json",
		"scene":  "",
		"path":   "",
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, _ := writer.CreateFormFile("file", fields["file"])

	_, _ = fw.Write(fileContent)
	for k, v := range fields {
		_ = writer.WriteField(k, v)
	}
	_ = writer.Close()

	resp, _ := http.Post(destURL, writer.FormDataContentType(), body)
	respbody, _ := ioutil.ReadAll(resp.Body)
	fileinfo := &params.FileInfo{}
	_ = json.Unmarshal(respbody, &fileinfo)

	return fileinfo.URL
}

func main() {

	defer exeDB.Close()
	path := os.Args[1]
	readImageFile(path)
	openFile()
	_, _ = ffmt.Puts(file_map)

}
