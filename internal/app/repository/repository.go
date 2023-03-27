package repository

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/repository/user"
)

var RepositorySet = wire.NewSet(
	user.UserRepositorySet,
)
