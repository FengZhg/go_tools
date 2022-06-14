package logger

import (
	"github.com/FengZhg/go_tools/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// @Author: Feng
// @Date: 2022/5/20 13:39

var (
	Log = NewLogger(zap.DebugLevel)
)

//WithLogLevel 更改日志等级
func WithLogLevel(level zapcore.Level) {
	Log = NewLogger(level)
}

func NewLogger(logLevel zapcore.Level) *zap.SugaredLogger {
	encoder := zapcore.NewConsoleEncoder(NewLoggerEncoderConfig())
	config := NewLoggerConfig(logLevel)
	// 构造cores
	syncers := buildWriteSyncers(&config)
	var cores []zapcore.Core
	for _, syncer := range syncers {
		cores = append(cores, zapcore.NewCore(encoder, syncer, zap.NewAtomicLevelAt(logLevel)))
	}
	logger, err := config.Build(
		zap.WrapCore(
			func(c zapcore.Core) zapcore.Core {
				return zapcore.NewTee(cores...)
			},
		),
	)
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}

func NewLoggerConfig(logLevel zapcore.Level) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    NewLoggerEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewLoggerEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        " ",
		LevelKey:       " ",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func buildWriteSyncers(config *zap.Config) (ret []zapcore.WriteSyncer) {
	// 加载输出路径
	for _, path := range config.OutputPaths {
		ws, _, err := zap.Open(path)
		if err != nil {
			panic(err)
		}
		ret = append(ret, ws)
	}
	// 加载文件
	ret = append(ret, zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/" + utils.GetExeFileName() + ".log", //Filename 是要写入日志的文件。
		MaxSize:    20,                                         //MaxSize 是日志文件在轮换之前的最大大小（以兆字节为单位）。它默认为 100 兆字节
		MaxBackups: 10,                                         //MaxBackups 是要保留的最大旧日志文件数。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
		MaxAge:     30,                                         //MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
		Compress:   true,                                       //压缩
		LocalTime:  true,                                       //LocalTime 确定用于格式化备份文件中的时间戳的时间是否是计算机的本地时间。默认是使用 UTC 时间。
	}))
	return ret
}
