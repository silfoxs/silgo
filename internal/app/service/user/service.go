package user

import (
	"github.com/google/wire"
	user_r "github.com/silfoxs/silgo/internal/app/repository/user"
)

var UserServiceSet = wire.NewSet(wire.Struct(new(UserService), "*"))

type UserService struct {
	UserRepo *user_r.UserRepository
}
