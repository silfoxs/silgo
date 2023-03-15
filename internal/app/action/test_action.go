package action

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Id   int32 `form:"id"`   // 主键ID
	Used int32 `form:"used"` // 是否启用 1:是 -1:否
}

type Response struct {
	Id int32 `json:"id"` // 主键ID
}

func (a *Action) Ping(c *gin.Context) {
	req := new(Request)
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "params is error",
		})
		return
	}
	params_str, _ := json.Marshal(req)
	a.logger.Info(string(params_str))
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"data": &Response{
			Id: req.Id,
		},
	})
}
