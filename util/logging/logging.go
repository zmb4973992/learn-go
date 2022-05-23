package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	//Logger, _ = zap.NewProduction() //zap的官方写法
	//defer Logger.Sync()             //zap的官方写法

	//
	//workingDirectory, _ := os.Getwd()
	//relativePath := util.MyLogConfig.RelativePath
	//Logger
	//Logger.Info("failed to fetch URL",
	//	// Structured context as strongly typed Field values.
	//	zap.String("url", "abc"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("name", time.Second),
	//)

}
