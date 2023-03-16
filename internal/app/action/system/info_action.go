package system

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
	Connection      string `json:"connection"` // 主键ID
	SecChUaPlatform string `json:"sec-ch-ua-platform"`
	UserAgent       string `json:"User-Agent"`
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
	a.logger.Info(string(params_str))
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"data": &Response{
			Connection:      c.GetHeader("connection"),
			SecChUaPlatform: c.GetHeader("sec-ch-ua-platform"),
			UserAgent:       c.GetHeader("User-Agent"),
		},
	})
}
