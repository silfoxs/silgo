package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/app/action/user"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"gorm.io/gorm"
)

type Options struct {
	Mode   string
	Logger *logger.Logger
	ReadDb *gorm.DB
}

func NewRouter(options Options) (*gin.Engine, error) {
	gin.SetMode(options.Mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	userAction := user.New(user.Options{
		Logger: options.Logger,
		ReadDb: options.ReadDb,
	})
	r.GET("/user/info", userAction.Info)

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
