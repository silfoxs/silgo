package user

import (
	"github.com/silfoxs/silgo/internal/app/service/user"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

type Options struct {
	Logger *logger.Logger
	ReadDb *gorm.DB
}

type Action struct {
	logger      *logger.Logger
	readDb      *gorm.DB
	userService *user.UserService
}

func New(options Options) *Action {
	return &Action{
		logger: options.Logger,
		readDb: options.ReadDb,
		userService: user.New(user.Options{
			Logger: options.Logger,
			ReadDb: options.ReadDb,
		}),
	}
}
