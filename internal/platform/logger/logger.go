package logger

import (
	"io"
	"os"
	"time"

	"github.com/Rocksus/fundtract/internal/model/constant"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	LogFormatLogFmt = "LOGFMT"
	LogFormatJSON   = "JSON"
)

// initLogFile initializes log file for the given path
func initLogFile(path, filename string) (*os.File, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(path+filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type LogManager struct {
	io.Writer
	ErrorWriter io.Writer
}

func (lm *LogManager) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	w := lm.Writer
	if level > zerolog.InfoLevel {
		w = lm.ErrorWriter
	}
	return w.Write(p)
}

// Init will initialize logging to file.
func Init(logLevel, logFormat string) error {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if logLevel == "TRACE" {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else if logLevel == "DEBUG" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	writers := []io.Writer{}

	if logFormat == LogFormatLogFmt {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
	}

	jsonWriters := []io.Writer{}
	jsonErrWriters := []io.Writer{}

	if logFormat == LogFormatJSON {
		jsonWriters = append(jsonWriters, os.Stdout)
		jsonErrWriters = append(jsonErrWriters, os.Stderr)
	}

	// save errors to path if any
	stdErrLoggerFile, err := initLogFile(".", constant.AppName+".stderr.log")
	if err != nil {
		return err
	}
	jsonErrWriters = append(jsonErrWriters, stdErrLoggerFile)

	writers = append(writers, &LogManager{
		Writer:      io.MultiWriter(jsonWriters...),
		ErrorWriter: io.MultiWriter(jsonErrWriters...),
	})

	multi := zerolog.MultiLevelWriter(
		writers...,
	)

	log.Logger = zerolog.
		New(multi).
		With().
		Stack().
		Timestamp().
		Logger()

	return nil
}
