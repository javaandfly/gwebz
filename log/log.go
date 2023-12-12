package log

import (
	"fmt"
	"io"

	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var logger = &lumberjack.Logger{
	Filename:   "",
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     28, //days
}

type LoggerLevel string

const (
	LogLevelPanic  LoggerLevel = "panic"
	LogLevelFatal  LoggerLevel = "fatal"
	LogLevelError  LoggerLevel = "error"
	LogLevelWarn   LoggerLevel = "warn"
	LogLevelInfo   LoggerLevel = "info"
	LogLevelDebug  LoggerLevel = "debug"
	LogLevelTrance LoggerLevel = "trance"
)

func InitLogrus(filePath, logName, logLevel string) {
	logFileName := filepath.Join(filePath, logName)
	logger.Filename = logFileName
	//logrus.JSONFormatter{} and logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	fileAndStdoutWriter := io.Writer(logger)
	logrus.SetOutput(fileAndStdoutWriter)
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("parse logs level err, %v, use defaule level 'info'", err)
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)

	logrus.Info("current log level is " + logLevel)
}
