package app

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	gin.SetMode(viper.GetString("app.mode"))
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	s := &http.Server{
		Addr:           ":" + viper.GetString("app.port"),
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
