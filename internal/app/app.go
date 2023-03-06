package app

import (
	"net/http"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	server := gin.New()
	gin.SetMode(viper.GetString("app.mode"))
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.GET("/ping", func(c *gin.Context) {
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
	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
