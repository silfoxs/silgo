package router

import (
	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/app/action"
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

	actions := action.New(action.Options{
		Logger: options.Logger,
	})
	r.GET("/ping", actions.Ping)
	return r, nil
}
