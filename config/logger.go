package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// 设置输出格式
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// 设置输出到文件
	logFile, err := os.OpenFile(
		filepath.Join("logs", "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err == nil {
		Logger.SetOutput(logFile)
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)
}

func GetLogger() *logrus.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
