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

func (l *Logger) Debug(args ...any) {
	l.logger.Debug(args)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) DPanic(args ...any) {
	l.logger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...any) {
	l.logger.DPanicf(template, args...)
}

func (l *Logger) Panic(args ...any) {
	l.logger.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.logger.Panicf(template, args...)
}

func (l *Logger) Fatal(args ...any) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.logger.Fatalf(template, args...)
}

func New(opt s_logger.Options) *Logger {
	logger := &Logger{}
	logger.Default(opt)
	return logger
}
