package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	ZapLogger   *zap.Logger
	SugarLogger *zap.SugaredLogger
)

func InitLogger() {

	logPath := ""
	name := "im"
	debug := true

	hook := lumberjack.Logger{
		Filename:   logPath, //日志文件路径
		MaxSize:    128,     // 每个日志文件保存的大小 单位：M
		MaxAge:     7,       //文件最多保存天数
		MaxBackups: 30,      //日志文件最多保存多少个备份
		Compress:   false,   // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	var writes []zapcore.WriteSyncer

	// 如果是开发环境，同时在控制台也输出
	if debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	} else {
		// 文件配置
		writes = append(writes, zapcore.AddSync(&hook))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("appName", name))

	// 构造日志
	ZapLogger = zap.New(core, caller, development, field)
	SugarLogger = ZapLogger.Sugar()
	ZapLogger.Info("ZapLogger 初始化成功")
}
