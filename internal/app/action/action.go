package action

import "github.com/silfoxs/silgo/internal/pkg/logger"

type Options struct {
	Logger *logger.Logger
}

type Action struct {
	logger *logger.Logger
}

func New(options Options) *Action {
	return &Action{
		logger: options.Logger,
	}
}
