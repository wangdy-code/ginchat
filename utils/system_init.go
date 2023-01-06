package utils

import (
	"fmt"
	"ginchat/common"
	"ginchat/config"
	"ginchat/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitConfig() error {
	common.VP = viper.GetViper()
	common.VP.SetConfigName("app")
	common.VP.AddConfigPath("config")
	return common.VP.ReadInConfig()

}

func InitMysql() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	},
	)
	mysqlConfig, err := config.GetMysqlConfig()
	if err != nil {
		panic("failed to read mysqlConfig")
	}

	common.DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		mysqlConfig.User, mysqlConfig.PassWord,
		mysqlConfig.Ip, mysqlConfig.Port, mysqlConfig.Database)),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	common.DB.AutoMigrate(&models.UserBasic{})
}
