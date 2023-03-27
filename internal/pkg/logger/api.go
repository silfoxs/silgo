package logger

import (
	s_logger "github.com/silfoxs/silgo/pkg/logger"
	"github.com/spf13/viper"
)

func New() (*Logger, func(), error) {
	logger := &Logger{}
	logger.Default(s_logger.Options{
		FileName:  viper.GetString("log.path"),
		Compress:  true,
		LocalTime: true,
	})
	clearup := func() {
	}
	return logger, clearup, nil
}
