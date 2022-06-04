package util

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type generalConfig struct {
	AppMode  string
	HttpPort string
}

type dbConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	Dsn        string // Data Source Name 数据库连接字符串
}

type jwtConfig struct {
	SecretKey []byte //这里不能用string，是jwt包的要求，否则报错
}

type logConfig struct {
	FileName  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
}

type uploadConfig struct {
	FullPath string
	MaxSize  int64
}

var (
	GeneralConfig = new(generalConfig)
	DBConfig      = new(dbConfig)
	JWTConfig     = new(jwtConfig)
	LogConfig     = new(logConfig)
	UploadConfig  = new(uploadConfig)
)

func LoadConfig() {
	config, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件路径错误：", err)
		return
	}
	GeneralConfig.AppMode = config.Section("general").Key("AppMode").MustString("debug")  //config中不填的话就默认为debug
	GeneralConfig.HttpPort = config.Section("general").Key("HttpPort").MustString("8000") //config中不填的话就默认为8000

	DBConfig.DbHost = config.Section("database").Key("DbHost").String()
	DBConfig.DbPort = config.Section("database").Key("DbPort").String()
	DBConfig.DbName = config.Section("database").Key("DbName").String()
	DBConfig.DbUsername = config.Section("database").Key("DbUser").String()
	DBConfig.DbPassword = config.Section("database").Key("DbPassword").String()
	DBConfig.Dsn = "sqlserver://" + DBConfig.DbUsername + ":" + DBConfig.DbPassword + "@" + DBConfig.DbHost + ":" + DBConfig.DbPort + "?database=" + DBConfig.DbName

	tempSecretKey := config.Section("jwt").Key("SecretKey").String() //配置里的密钥是string类型，jwt要求为[]byte类型，必须转换后才能使用，否则就为空
	JWTConfig.SecretKey = []byte(tempSecretKey)

	LogConfig.FileName = config.Section("log").Key("FileName").MustString("D:/test/log/status.log")
	LogConfig.MaxSize = config.Section("log").Key("MaxSize").MustInt(1)
	LogConfig.MaxBackup = config.Section("log").Key("MaxBackup").MustInt(1)
	LogConfig.MaxAge = config.Section("log").Key("MaxAge").MustInt(1)

	UploadConfig.FullPath = config.Section("upload_files").Key("FullPath").MustString("D:/test/upload_files") + "/"
	UploadConfig.MaxSize = config.Section("upload_files").Key("MaxSize").MustInt64(10 << 20)
}
