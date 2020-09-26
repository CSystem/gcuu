package mlog

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
	file *os.File
}

func (l logger) Close() error {
	return l.file.Close()
}

var Logger *logger

// Setup logger
func Setup(env string) {
	log := logrus.New()                                                         // 实例化
	fileName := fmt.Sprintf("logs/%s.log", env)                                 // 日志文件
	src, err := os.OpenFile(os.DevNull, os.O_WRONLY|os.O_APPEND, os.ModeAppend) // 禁止 logrus 的输出

	if err != nil {
		// log.Info("Failed to log to file, using default stderr, err", err)
		fmt.Fprintf(os.Stderr, "failed to initialize log file %s", err)
		os.Exit(1)
	}

	log.Out = src                   // 设置输出
	log.SetLevel(logrus.DebugLevel) // 设置日志级别

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",                    // 分割后的文件名称
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 设置最大保存时间(7天)
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)

	if err != nil {
		// log.Info("Failed to rotatelogs to log file, err", err)
		fmt.Fprintf(os.Stderr, "failed to rotatelogs to log file %s", err)
		os.Exit(1)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	log.AddHook(lfHook)

	Logger = &logger{
		Logger: log,
		file:   src,
	}
}
