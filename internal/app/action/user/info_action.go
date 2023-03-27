package user

import (
	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/pkg/response"
)

type Request struct {
	Id int32 `form:"id" binding:"required,number"` // 用户ID
}

type Response struct {
	Id       int32  `json:"id"`        // 用户ID
	NickName string `json:"nick_name"` // 用户昵称
}

func (a *Action) Info(c *gin.Context) {
	req := new(Request)
	if err := c.ShouldBindQuery(req); err != nil {
		response.Fail(c, map[string]any{
			"msg": err.Error(),
		})
		return
	}
	user, err := a.UserService.Info(req.Id)
	if err != nil {
		response.Fail(c, map[string]any{
			"msg": err.Error(),
		})
		return
	}
	a.Logger.Infof("%+v%+v", req, user)
	response.Success(c, map[string]any{
		"data": user,
	})
}
