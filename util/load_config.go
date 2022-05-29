package util

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type DBConfig struct {
	AppMode    string
	HttpPort   string
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	Dsn        string // Data Source Name 数据库连接字符串
}

type JWTConfig struct {
	SecretKey []byte //这里不能用string，是jwt包的要求，否则报错
}

type LogConfig struct {
	RelativePath string
}

type UploadConfig struct {
	FullPath string
}

// MyConfig 结构体实例化
var (
	MyConfig       = new(DBConfig)
	MyJWTConfig    = new(JWTConfig)
	MyLogConfig    = new(LogConfig)
	MyUploadConfig = new(UploadConfig)
)

func LoadConfig() {
	config, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件路径错误：", err)
		return
	}
	MyConfig.AppMode = config.Section("server").Key("AppMode").MustString("debug")  //config中不填的话就默认为debug
	MyConfig.HttpPort = config.Section("server").Key("HttpPort").MustString("8000") //config中不填的话就默认为8000
	MyConfig.DbHost = config.Section("database").Key("DbHost").String()
	MyConfig.DbPort = config.Section("database").Key("DbPort").String()
	MyConfig.DbName = config.Section("database").Key("DbName").String()
	MyConfig.DbUsername = config.Section("database").Key("DbUser").String()
	MyConfig.DbPassword = config.Section("database").Key("DbPassword").String()
	MyConfig.Dsn = "sqlserver://" + MyConfig.DbUsername + ":" + MyConfig.DbPassword + "@" + MyConfig.DbHost + ":" + MyConfig.DbPort + "?database=" + MyConfig.DbName

	tempSecretKey := config.Section("jwt").Key("SecretKey").String() //配置里的密钥是string类型，jwt要求为[]byte类型，必须转换后才能使用，否则就为空
	MyJWTConfig.SecretKey = []byte(tempSecretKey)

	err = config.Section("log").MapTo(MyLogConfig)
	if err != nil {
		fmt.Println(err)
	}

	MyUploadConfig.FullPath = config.Section("upload_files").Key("FullPath").MustString("D:/test/upload_files") + "/" //config中不填的话会有默认值

}
