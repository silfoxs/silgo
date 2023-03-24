package logger

import (
	"github.com/google/wire"
	s_logger "github.com/silfoxs/silgo/pkg/logger"
	"github.com/spf13/viper"
)

var LoggerSet = wire.NewSet(New)

func New() (*Logger, func(), error) {
	logger := &Logger{}
	logger.Default(s_logger.Options{
		FileName:  viper.GetString("log.path"),
		Compress:  true,
		LocalTime: true,
	})
	return logger, nil, nil
}
