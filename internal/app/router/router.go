package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/silfoxs/silgo/internal/app/action/user"
	"github.com/spf13/viper"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	UserAction user.Action
}

func (r *Router) GetEngine() (*gin.Engine, error) {
	gin.SetMode(viper.GetString("app.mode"))
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	return router, nil
}

func (r *Router) Handler() (*gin.Engine, error) {
	router, err := r.GetEngine()
	if err != nil {
		return nil, err
	}
	router.GET("/user/info", r.UserAction.Info)

	router.NoMethod(r.noMethod)
	router.NoRoute(r.noRoute)
	return router, nil
}

func (r *Router) noMethod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  "request not allow method",
	})
}

func (r *Router) noRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  "request uri not exists",
	})
}
