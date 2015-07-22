package utils

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var Logger *log.Logger

func init() {
	var loggerSettings struct{ LogFilePath string }
	err := LoadSetting(&loggerSettings)

	if err != nil {
		panic(err.Error())
	}

	Logger = log.New(&lumberjack.Logger{
		Filename:   loggerSettings.LogFilePath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}, "", log.Ldate|log.Ltime|log.LstdFlags)

	Logger.Print("Init logger\n")
	Logger.Printf("Writing log to %s", loggerSettings.LogFilePath)
}
