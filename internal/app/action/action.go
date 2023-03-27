package action

import (
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/action/user"
)

var ActionSet = wire.NewSet(
	user.UserActionSet,
)
