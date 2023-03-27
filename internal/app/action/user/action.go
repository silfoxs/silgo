package user

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/service/user"
	"github.com/silfoxs/silgo/internal/pkg/logger"
)

var UserActionSet = wire.NewSet(wire.Struct(new(Action), "*"))

type Action struct {
	Logger      *logger.Logger
	UserService *user.UserService
}
