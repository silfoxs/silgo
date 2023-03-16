package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Id int32 `form:"id"` // 用户ID
}

type Response struct {
	Id       int32  `json:"id"`        // 用户ID
	NickName string `json:"nick_name"` // 用户昵称
}

func (a *Action) Info(c *gin.Context) {
	req := new(Request)
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "params is error",
		})
		return
	}
	params_str, _ := json.Marshal(req)
	user := a.userService.Info(req.Id)
	a.logger.Info(string(params_str), user)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "",
		"data": user,
	})
}
