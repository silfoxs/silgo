package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type Logger struct {
	opt Options
}

func (l *Logger) SetOptions(opt Options) {
	l.opt = opt
}

func (l *Logger) GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *Logger) JsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
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
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func New(opt Options) *zap.SugaredLogger {
	logger := &Logger{opt: opt}
	return zap.New(
		zapcore.NewCore(logger.GetEncoder(), logger.GetLogWriter(), zapcore.DebugLevel),
		zap.AddCaller(),
	).Sugar()
}

func NewJson(opt Options) *zap.SugaredLogger {
	logger := &Logger{opt: opt}
	return zap.New(
		zapcore.NewCore(logger.JsonEncoder(), logger.GetLogWriter(), zapcore.DebugLevel),
		zap.AddCaller(),
	).Sugar()
}
