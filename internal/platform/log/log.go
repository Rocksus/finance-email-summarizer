package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger map[LogLevel]*logrus.Logger

type LogLevel int

const (
	TraceLevel = LogLevel(logrus.TraceLevel)
	DebugLevel = LogLevel(logrus.DebugLevel)
	InfoLevel  = LogLevel(logrus.InfoLevel)
	WarnLevel  = LogLevel(logrus.WarnLevel)
	ErrorLevel = LogLevel(logrus.ErrorLevel)
	FatalLevel = LogLevel(logrus.FatalLevel)
)

func initLogFile(path string, lg *logrus.Logger, level LogLevel) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	lg.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	lg.SetLevel(logrus.Level(level))
	lg.SetOutput(f)

	return nil
}

// InitLogger will initialize logging to file
func Init(path string) error {
	traceLogger := logrus.New()
	debugLogger := logrus.New()
	infoLogger := logrus.New()
	warnLogger := logrus.New()
	errorLogger := logrus.New()
	fatalLogger := logrus.New()

	logger = make(map[LogLevel]*logrus.Logger)

	logger[TraceLevel] = traceLogger
	logger[DebugLevel] = debugLogger
	logger[InfoLevel] = infoLogger
	logger[WarnLevel] = warnLogger
	logger[ErrorLevel] = errorLogger
	logger[FatalLevel] = fatalLogger

	if path == "" {
		return nil
	}

	if err := initLogFile(path+".debug.log", debugLogger, DebugLevel); err != nil {
		return err
	}

	if err := initLogFile(path+".info.log", infoLogger, InfoLevel); err != nil {
		return err
	}

	if err := initLogFile(path+".warn.log", warnLogger, WarnLevel); err != nil {
		return err
	}

	if err := initLogFile(path+".error.log", errorLogger, ErrorLevel); err != nil {
		return err
	}

	if err := initLogFile(path+".fatal.log", fatalLogger, FatalLevel); err != nil {
		return err
	}

	return nil
}

func Error(args ...interface{}) {
	if logger[ErrorLevel] == nil {
		fmt.Println(args...)
		return
	}
	logger[ErrorLevel].Error(args...)
}

func ErrorWithFields(msg string, fields KV) {
	if logger[ErrorLevel] == nil {
		fmt.Println(msg, fields)
		return
	}
	logger[ErrorLevel].WithFields(logrus.Fields(fields)).Error(msg)
}

func Info(args ...interface{}) {
	if logger[InfoLevel] == nil {
		fmt.Println(args...)
		return
	}
	logger[InfoLevel].Info(args...)
}

func InfoWithFields(msg string, fields KV) {
	if logger[InfoLevel] == nil {
		fmt.Println(msg, fields)
		return
	}
	logger[InfoLevel].WithFields(logrus.Fields(fields)).Info(msg)
}

func Debug(args ...interface{}) {
	if logger[DebugLevel] == nil {
		fmt.Println(args...)
		return
	}
	logger[DebugLevel].Debug(args...)
}

func DebugWithFields(msg string, fields KV) {
	if logger[DebugLevel] == nil {
		fmt.Println(msg, fields)
		return
	}
	logger[DebugLevel].WithFields(logrus.Fields(fields)).Debug(msg)
}

func Warn(args ...interface{}) {
	if logger[WarnLevel] == nil {
		fmt.Println(args...)
		return
	}
	logger[WarnLevel].Warn(args...)
}

func WarnWithFields(msg string, fields KV) {
	if logger[WarnLevel] == nil {
		fmt.Println(msg, fields)
		return
	}
	logger[WarnLevel].WithFields(logrus.Fields(fields)).Warn(msg)
}

func Fatal(args ...interface{}) {
	if logger[FatalLevel] == nil {
		fmt.Println(args...)
		os.Exit(0)
	}
	logger[FatalLevel].Fatal(args...)
}

func FatalWithFields(msg string, fields KV) {
	if logger[FatalLevel] == nil {
		fmt.Println(msg, fields)
		os.Exit(0)
	}
	logger[FatalLevel].WithFields(logrus.Fields(fields)).Fatal(msg)
}
