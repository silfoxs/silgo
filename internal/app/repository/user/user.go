package user

import (
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

type Options struct {
	Logger *logger.Logger
	ReadDb *gorm.DB
}

type UserRepository struct {
	logger *logger.Logger
	readDb *gorm.DB
}

type User struct {
	Id       int32  `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
}

func (u *UserRepository) Info(id int32) (User, error) {
	var ret User
	if err := u.readDb.Where("id=?", id).First(&ret).Error; err != nil {
		u.logger.Errorf("user id %d info error: %s", id, err.Error())
		return ret, err
	}
	return ret, nil
}

func New(options Options) *UserRepository {
	return &UserRepository{
		logger: options.Logger,
		readDb: options.ReadDb,
	}
}
