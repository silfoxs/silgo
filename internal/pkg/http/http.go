package http

import (
	"net/http"

	"time"

	"fmt"

	"context"

	"os"

	"github.com/silfoxs/silgo/internal/app/router"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"github.com/spf13/viper"
)

func NewHttp(log *logger.Logger, router *router.Router) (*http.Server, func(), error) {
	server, err := router.Handler()
	if err != nil {
		return nil, nil, err
	}
	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler:        server,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	clearup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		pid := os.Getpid()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Errorf("pid:%d server close err %s", pid, err.Error())
		} else {
			log.Infof("pid:%d server close success", pid)
		}
	}
	return httpServer, clearup, nil
}
