package app

import (
	"context"
	"net/http"
	"os"

	"time"

	"fmt"

	"github.com/silfoxs/silgo/internal/app/router"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	"github.com/silfoxs/silgo/pkg/shutdown"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	log    *logger.Logger
	readDb *gorm.DB
)

func Run() {
	pid := os.Getpid()
	_, err := initServer()
	if err != nil {
		log.Panicf("pid:%d init server error: %s", pid, err.Error())
	}
	server, err := router.NewRouter(router.Options{
		Mode:   viper.GetString("app.mode"),
		Logger: log,
		ReadDb: readDb,
	})
	if err != nil {
		log.Panicf("pid:%d server start error: %s", pid, err.Error())
	}

	log.Infof("pid:%d server start", pid)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Errorf("pid:%d %s", pid, err.Error())
		}
	}()
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := s.Shutdown(ctx); err != nil {
				log.Errorf("pid:%d server close err %s", pid, err.Error())
			} else {
				log.Infof("pid:%d server close success", pid)
			}
		},
	)
}

func initServer() (func(), error) {
	injector, _, err := BuildInjector()
	if err != nil {
		return nil, err
	}
	log = injector.Logger
	readDb = injector.Db
	return nil, nil
}
