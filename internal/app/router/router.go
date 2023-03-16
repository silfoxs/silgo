package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/app/action/system"
	"github.com/silfoxs/silgo/internal/pkg/logger"
)

type Options struct {
	Mode   string
	Logger *logger.Logger
}

func NewRouter(options Options) (*gin.Engine, error) {
	gin.SetMode(options.Mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	system := system.New(system.Options{
		Logger: options.Logger,
	})
	r.GET("/system/info", system.Info)

	r.NoMethod(noMethod)
	r.NoRoute(noRoute)
	return r, nil
}

func noMethod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  "request not allow method",
	})
}

func noRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  "request uri not exists",
	})
}
