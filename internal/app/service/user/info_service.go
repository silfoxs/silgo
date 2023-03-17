package user

import "github.com/silfoxs/silgo/internal/app/repository/user"

func (u *UserService) Info(id int32) (user.User, error) {
	return u.userRep.Info(id)
}
