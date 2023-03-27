package database

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func New() (*gorm.DB, func(), error) {
	db := &Mysql{
		Host:     viper.GetString("mysql.read.host"),
		Port:     viper.GetInt("mysql.read.port"),
		UserName: viper.GetString("mysql.read.username"),
		Password: viper.GetString("mysql.read.password"),
		Database: viper.GetString("mysql.read.database"),
	}
	con, err := db.New()
	if err != nil {
		return nil, nil, err
	}
	return con, nil, nil
}
