package app

import (
	"context"
	"encoding/json"
	"net/http"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/pkg/logger"
	s_log "github.com/silfoxs/silgo/pkg/logger"
	"github.com/silfoxs/silgo/pkg/shutdown"
	"github.com/spf13/viper"
)

func Run() {
	gin.SetMode(viper.GetString("app.mode"))
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.GET("/ping", func(c *gin.Context) {
		log := logger.New(s_log.Options{
			FileName: viper.GetString("log.path"),
		})
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
	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := s.Shutdown(ctx); err != nil {
				fmt.Println("server shutdown err" + err.Error())
			} else {
				fmt.Println("server shutdown success")
			}
		},
	)
}
