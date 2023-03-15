package logger

import (
	s_logger "github.com/silfoxs/silgo/pkg/logger"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func (l *Logger) Default(opt s_logger.Options) {
	l.logger = s_logger.NewJson(opt)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func New(opt s_logger.Options) *Logger {
	logger := &Logger{}
	logger.Default(opt)
	return logger
}
