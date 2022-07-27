package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	appConfig
	dbConfig
	jwtConfig
	logConfig
	uploadConfig
	roleConfig
}

type appConfig struct {
	AppMode  string
	HttpPort string
}

type dbConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	DSN        string // Data Source Name 数据库连接字符串
}

type jwtConfig struct {
	SecretKey      []byte //这里不能用string，是jwt包的要求，否则报错
	ValidityPeriod int
}

type logConfig struct {
	Path          string
	FileName      string
	MaxSizeForLog int
	MaxBackup     int
	MaxAge        int
	Compress      bool
}

type uploadConfig struct {
	FullPath         string
	MaxSizeForUpload int64
}

type roleConfig struct {
	level string
	name  string
}

var (
	GlobalConfig = new(config)
	v            = viper.New()
)

func Init() {
	v.SetConfigName("config")
	v.SetConfigType("ini")
	v.AddConfigPath("./config/")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file error: %w \n", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已修改:", e.Name)
		loadConfig()
	})
	v.WatchConfig()
	loadConfig()
}

func loadConfig() {
	GlobalConfig.appConfig.AppMode = v.GetString("app.AppMode")
	GlobalConfig.appConfig.HttpPort = v.GetString("app.HttpPort")

	GlobalConfig.dbConfig.DbHost = v.GetString("database.DbHost")
	GlobalConfig.dbConfig.DbPort = v.GetString("database.DbPort")
	GlobalConfig.dbConfig.DbName = v.GetString("database.DbName")
	GlobalConfig.dbConfig.DbUsername = v.GetString("database.DbUsername")
	GlobalConfig.dbConfig.DbPassword = v.GetString("database.DbPassword")
	GlobalConfig.dbConfig.DSN =
		"sqlserver://" + GlobalConfig.dbConfig.DbUsername + ":" +
			GlobalConfig.dbConfig.DbPassword + "@" + GlobalConfig.dbConfig.DbHost +
			":" + GlobalConfig.dbConfig.DbPort + "?database=" + GlobalConfig.dbConfig.DbName

	//配置里的密钥是string类型，jwt要求为[]byte类型，必须转换后才能使用
	GlobalConfig.jwtConfig.SecretKey = []byte(v.GetString("jwt.SecretKey"))
	GlobalConfig.jwtConfig.ValidityPeriod = v.GetInt("jwt.ValidityPeriod")

	GlobalConfig.logConfig.Path = v.GetString("log.LogPath")
	GlobalConfig.logConfig.FileName = v.GetString("log.LogPath") + "/status.log"
	GlobalConfig.logConfig.MaxSizeForLog = v.GetInt("log.LogMaxSize")
	GlobalConfig.logConfig.MaxBackup = v.GetInt("log.LogMaxBackup")
	GlobalConfig.logConfig.MaxAge = v.GetInt("log.LogMaxAge")
	GlobalConfig.logConfig.Compress = v.GetBool("log.LogCompress")

	GlobalConfig.uploadConfig.FullPath = v.GetString("upload_files.FullPath") + "/"
	GlobalConfig.uploadConfig.MaxSizeForUpload = v.GetInt64("upload_files.MaxSize") << 20

}
