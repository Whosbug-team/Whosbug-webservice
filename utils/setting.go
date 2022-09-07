package utils

// 对ini配置文件进行读取和处理

import (
	"gopkg.in/ini.v1"
	"log"
)

// 变量引入
var (
	AppMode  string
	HTTPPort string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSSLMode  string
	DbTimeZone string
)

func init() {
	file, err := ini.Load("configs/config.ini")
	if err != nil {
		log.Fatalf("ini配置文件读取出错，请检查文件路径: ", err)
	}
	SrvSection := file.Section("server")
	LoadServer(SrvSection)
	DBSection := file.Section("database")
	LoadDataBase(DBSection)
}

// LoadServer 加载服务
func LoadServer(section *ini.Section) {
	AppMode = section.Key("AppMode").MustString("debug") // 默认设为debug模式
	HTTPPort = section.Key("HttpPort").MustString(":8083")
}

// LoadDataBase 加载数据库
func LoadDataBase(section *ini.Section) {
	DbHost = section.Key("DbHost").MustString("localhost")
	DbPort = section.Key("DbPort").MustString("5432")
	DbUser = section.Key("DbUser").MustString("postgres")
	DbPassword = section.Key("DbPassword").MustString("123456")
	DbName = section.Key("DbName").MustString("whobug2022")
	DbSSLMode = section.Key("DbSSLMode").MustString("disable")
	DbTimeZone = section.Key("DbTimeZone").MustString("PRC")
}
