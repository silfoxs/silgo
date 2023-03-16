package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LocalTime  bool
	Compress   bool
}

type Logger struct {
	opt Options
}

func (l *Logger) SetOptions(opt Options) {
	l.opt = opt
}

func (l *Logger) GetEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return encoderConfig
}

func (l *Logger) GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(l.GetEncoderConfig())
}

func (l *Logger) JsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(l.GetEncoderConfig())
}

func (l *Logger) GetLogWriter() zapcore.WriteSyncer {
	dir := filepath.Dir(l.opt.FileName)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   l.opt.FileName,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  l.opt.LocalTime,
		Compress:   l.opt.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func New(opt Options) *zap.SugaredLogger {
	logger := &Logger{opt: opt}
	return zap.New(
		zapcore.NewCore(logger.GetEncoder(), logger.GetLogWriter(), zapcore.DebugLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	).Sugar()
}

func NewJson(opt Options) *zap.SugaredLogger {
	logger := &Logger{opt: opt}
	return zap.New(
		zapcore.NewCore(logger.JsonEncoder(), logger.GetLogWriter(), zapcore.DebugLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	).Sugar()
}
