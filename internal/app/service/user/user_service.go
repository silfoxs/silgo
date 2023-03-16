package user

import (
	user_r "github.com/silfoxs/silgo/internal/app/repository/user"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

type Options struct {
	Logger *logger.Logger
	ReadDb *gorm.DB
}

type UserService struct {
	logger  *logger.Logger
	readDb  *gorm.DB
	userRep *user_r.UserRepository
}

func New(options Options) *UserService {
	return &UserService{
		logger: options.Logger,
		readDb: options.ReadDb,
		userRep: user_r.New(user_r.Options{
			Logger: options.Logger,
			ReadDb: options.ReadDb,
		}),
	}
}
