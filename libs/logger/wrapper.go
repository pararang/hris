package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Warn(msg string)
	Error(msg string, err error)
}

type logger struct {
	writer *zap.Logger
}

func New() Logger {
	return &logger{
		writer: zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), os.Stdout, zap.InfoLevel)),
	}
}

func (l *logger) Warn(msg string) {
	l.writer.Error(msg)
}

func (l *logger) Error(msg string, err error) {
	l.writer.Error(msg, zap.Error(err))
}
