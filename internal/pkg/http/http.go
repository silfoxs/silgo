package http

import (
	"net/http"

	"time"

	"fmt"

	"github.com/silfoxs/silgo/internal/app/router"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewHttp(log *logger.Logger, db *gorm.DB) (*http.Server, error) {
	server, err := router.NewRouter(router.Options{
		Mode:   viper.GetString("app.mode"),
		Logger: log,
		ReadDb: db,
	})
	if err != nil {
		return nil, err
	}
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, nil
}
