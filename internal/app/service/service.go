package service

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/service/user"
)

var ServiceSet = wire.NewSet(
	user.UserServiceSet,
)
