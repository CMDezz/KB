package logger

import (
	"io"

	"github.com/CMDezz/KB/utils/constants"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var mLog *log.Logger

func NewLogger(path string, enviroment string) *log.Logger {

	if enviroment == "dev" {
		mLog = log.New()
		mLog.Formatter = &log.TextFormatter{TimestampFormat: constants.TimeFormatDefault}
	} else if enviroment == "prod" {
		mLog = log.New()
		mLog.Out = io.Discard
		mLog.Hooks.Add(lfshook.NewHook(
			lfshook.PathMap{
				log.InfoLevel:  path + "/" + enviroment + "_info.log",
				log.TraceLevel: path + "/" + enviroment + "_trace.log",
				log.WarnLevel:  path + "/" + enviroment + "_warn.log",
				log.DebugLevel: path + "/" + enviroment + "_debug.log",
				log.ErrorLevel: path + "/" + enviroment + "_error.log",
				log.FatalLevel: path + "/" + enviroment + "_fatal.log",
				log.PanicLevel: path + "/" + enviroment + "_panic.log",
			},
			&log.JSONFormatter{
				TimestampFormat: constants.TimeFormatDefault,
			},
		))
	} else {
		Error("Error init Logger")
	}
	return mLog
}

func Debug(format string, v ...any) {
	mLog.Debugf(constants.LogDebugPrefix+format, v...)
}

func Info(format string, v ...any) {
	mLog.Infof(constants.LogInfoPrefix+format, v...)
}

func Warn(format string, v ...any) {
	mLog.Warnf(constants.LogWarnPrefix+format, v...)
}

func Error(format string, v ...any) {
	mLog.Errorf(constants.LogErrorPrefix+format, v...)
}

func Fatal(format string, v ...any) {
	mLog.Fatalf(constants.LogFatalPrefix+format, v...)
}

func Panic(format string, v ...any) {
	mLog.Panicf(constants.LogPanicPrefix+format, v...)
}
