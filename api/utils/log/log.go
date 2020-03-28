package log

import (
	logger "github.com/sirupsen/logrus"
)

func Setup() {
	logger.SetLevel(logger.DebugLevel)
	logger.SetFormatter(&logger.TextFormatter{
		ForceColors: true,
	}) //JSONFormatter
	//logger.SetReportCaller(true)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
