// Package v1 处理业务逻辑, goRJ 控制器 v1
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
	"time"
)

// BaseAPIController 基础控制器
type BaseApiController struct {
}

const (
	SUCCESS int = 200 //操作成功
	FAILED  int = 500 //操作失败
)

// 请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    SUCCESS,
		"message": "成功!",
		"data":    v,
	})
}

// 请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    FAILED,
		"data":    nil,
		"message": v,
	})
}

func setupLogger() *zap.Logger {
	logPath := "E:/www-go/RuiJi/logs"
	isFile(logPath)
	currentTime := time.Now()
	timeStr := currentTime.Format("2006-01-02")
	// 日志文件的输出路径和文件名
	logFile := logPath + "/gin-" + timeStr + ".log"
	// 创建Lumberjack实例，用于日志文件的切割和清理
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100,   // 单个日志文件最大大小，单位为MB
		MaxBackups: 10,    // 最多保留的日志文件数
		MaxAge:     30,    // 日志文件最长保留时间，单位为天
		Compress:   false, // 是否启用压缩
	}

	// 配置Zap日志库
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(lumberjackLogger),
		zap.NewAtomicLevel(),
	)

	return zap.New(core, zap.AddCaller())
}

// 判断日志存放文件夹是否存在
func isFile(dirPath string) {
	// 检查文件夹是否存在
	_, err := os.Stat(dirPath)
	// 如果文件夹不存在，则创建
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			// 处理创建文件夹失败的情况
			panic(err)
		}
	}
}
