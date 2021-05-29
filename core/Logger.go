package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Logger struct {
	logger *zap.Logger
}

var onceMutex = sync.Once{}
var logger *Logger

func GetLogger() *Logger {
	onceMutex.Do(func() {
		logger = &Logger{}
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
			EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		}
		atom := zap.NewAtomicLevelAt(zap.DebugLevel)
		logConfig := zap.Config{
			Level:            atom,                                                // 日志级别
			Development:      false,                                               // 开发模式，堆栈跟踪
			Encoding:         "json",                                              // 输出格式 console 或 json
			EncoderConfig:    encoderConfig,                                       // 编码器配置
			InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
			ErrorOutputPaths: []string{"stderr"},
			OutputPaths:      []string{"stdout"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		}
		log, _ := logConfig.Build()
		logger.logger = log
	})
	return logger
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}
