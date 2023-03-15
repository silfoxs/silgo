package app

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	s_log "github.com/silfoxs/silgo/pkg/logger"
	"github.com/silfoxs/silgo/pkg/shutdown"
	"github.com/spf13/viper"
)

func Run() {
	log := logger.New(s_log.Options{
		FileName: viper.GetString("log.path"),
	})
	gin.SetMode(viper.GetString("app.mode"))
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.GET("/ping", func(c *gin.Context) {
		json_str, _ := json.Marshal(map[string]interface{}{
			"path": viper.AllSettings(),
		})
		log.Info(string(json_str))
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infof("pid:%d server start", os.Getpid())
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Errorf("pid:%d %s", os.Getpid(), err.Error())
		}
	}()
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := s.Shutdown(ctx); err != nil {
				log.Errorf("pid:%d server close err %s", err.Error(), os.Getpid())
			} else {
				log.Infof("pid:%d server close success", os.Getpid())
			}
		},
	)
}
