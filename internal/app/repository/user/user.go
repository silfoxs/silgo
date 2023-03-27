package user

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

var UserRepositorySet = wire.NewSet(wire.Struct(new(UserRepository), "*"))

type UserRepository struct {
	Logger *logger.Logger
	ReadDb *gorm.DB
}

type User struct {
	Id       int32  `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
}

func (u *UserRepository) Info(id int32) (User, error) {
	var ret User
	if err := u.ReadDb.Where("id=?", id).First(&ret).Error; err != nil {
		u.Logger.Errorf("user id %d info error: %s", id, err.Error())
		return ret, err
	}
	return ret, nil
}
