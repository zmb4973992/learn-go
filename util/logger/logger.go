package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"learn-go/util"
	"os"
	"time"
)

var (
	ZapStandardLogger *zap.Logger        //zap的标准logger，速度更快，但是输入麻烦，用于取代gin的logger
	L                 *zap.SugaredLogger // zap的加糖logger，速度慢一点点，但是输入方便，自己用
)

func InitLogger() {
	encoder := newEncoder() //调用自定义的编码器函数，生成新的编码器
	//调用自定义的写入同步器函数，传入文件路径+名称、最大尺寸、最大备份数量、最大保存天数，生成新的写入同步器
	writeSyncer := newWriteSyncer(util.LogConfig.FileName, util.LogConfig.MaxSize, util.LogConfig.MaxBackup, util.LogConfig.MaxAge)
	mode := util.GeneralConfig.AppMode
	//声明zap的核心参数
	var core zapcore.Core
	//如果是开发模式：
	if mode == "debug" || mode == "dev" {
		developmentEncoderConfig := zap.NewDevelopmentEncoderConfig()                            //生成开发模式下的、默认encoder配置文件，用于管理控制台的显示内容
		developmentEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") //调整格式
		consoleEncoder := zapcore.NewConsoleEncoder(developmentEncoderConfig)                    //根据encoder配置文件，生成控制台的encoder
		core = zapcore.NewTee(                                                                   //生成zap的核心文件core
			zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),                    // 往日志文件里面写
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 在终端输出
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) //如果是生产环境就只写到日志里，不在终端输出
	}
	ZapStandardLogger = zap.New(core, zap.AddCaller()) //根据zap的要求，生成一个日志记录器
	L = ZapStandardLogger.Sugar()                      //使用加糖模式的日志记录器，牺牲点效率，但简单一些
	defer L.Sync()
}

// 自定义 生成写入同步器的函数
func newWriteSyncer(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  //日志文件的位置
		MaxSize:    maxSize,   //在进行切割之前，日志文件的最大大小（MB）
		MaxBackups: maxBackup, //保留旧文件的最大个数
		MaxAge:     maxAge,    //保留旧文件的最大天数
		Compress:   false,     //是否压缩旧文件，默认为否
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 自定义 生成编码器的函数
func newEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// ZapLogger 接管gin框架默认的日志，用作中间件
func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
