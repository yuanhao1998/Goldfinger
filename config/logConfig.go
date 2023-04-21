// @Create   : 2023/4/19 10:28
// @Author   : yaho
// @Remark   : 日志配置

package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

func InitLog(level string, path string) *zap.SugaredLogger {

	levelMap := map[string]zapcore.Level{
		"debug":   zapcore.DebugLevel,
		"info":    zapcore.InfoLevel,
		"warn":    zapcore.WarnLevel,
		"warning": zapcore.WarnLevel,
		"error":   zapcore.ErrorLevel,
		"dPanic":  zapcore.DPanicLevel,
		"panic":   zapcore.PanicLevel,
		"fatal":   zapcore.FatalLevel,
	}

	logDirPath, err := filepath.Abs("./log/")
	if err != nil {
		panic("读取日志目录绝对路径失败：" + err.Error())
	}

	encoder := setupLogFormat()
	writeSyncer := getFileLogWriter(logDirPath + path)
	core := zapcore.NewTee(
		// 同时向控制台和文件写日志
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), levelMap[level]),
		zapcore.NewCore(encoder, writeSyncer, levelMap[level]),
	)

	return zap.New(core, zap.AddCaller()).Sugar()

}

// 配置日志格式
func setupLogFormat() zapcore.Encoder {

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(logTmFmtWithMS) + "]")
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder

	return zapcore.NewConsoleEncoder(encoderConfig) // NewConsoleEncoder 打印更符合人们观察的方式

}

// 切分日志
func getFileLogWriter(path string) (writeSyncer zapcore.WriteSyncer) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    LogMaxSize,
		MaxBackups: LogFileMaxNum,
		MaxAge:     LogFileSplitDay,
		Compress:   true,
	}

	return zapcore.AddSync(lumberJackLogger)
}
